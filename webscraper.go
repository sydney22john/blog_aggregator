package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sjohn/blog_aggregator/internal/database"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Rss xml structure
type Rss struct {
	XMLName xml.Name `xml:"rss"`
	// Text    string   `xml:",chardata"`
	// Atom    string   `xml:"atom,attr"`
	// Version string   `xml:"version,attr"`
	Channel struct {
		// Text  string `xml:",chardata"`
		Title string `xml:"title"`
		// Link  struct {
		// 	Text string `xml:",chardata"`
		// 	Href string `xml:"href,attr"`
		// 	Rel  string `xml:"rel,attr"`
		// 	Type string `xml:"type,attr"`
		// } `xml:"link"`
		// Description string `xml:"description"`
		// Generator     string `xml:"generator"`
		// Language      string `xml:"language"`
		// LastBuildDate string `xml:"lastBuildDate"`
		Item []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Text        string `xml:",chardata"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
	Description string `xml:"description"`
}

func getRssData(url string) (Rss, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Rss{}, err
	}

	rss := Rss{}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Rss{}, err
	}
	if err := xml.Unmarshal(respBytes, &rss); err != nil {
		return Rss{}, err
	}
	return rss, nil
}

func (cfg *apiConfig) blogAggregatorWorker(processFeeds int) {
	var wg sync.WaitGroup
	errChan := make(chan error, processFeeds)

	for {
		log.Println("Worker: getting top 10 feeds from db")
		feeds, err := cfg.DB.GetNextFeedsToFetch(context.Background(), int32(processFeeds))
		if err != nil {
			log.Println(err)
		}

		log.Println("Worker: starting go routines")
		for _, feed := range feeds {
			wg.Add(1)
			go cfg.aggBlog(&wg, feed, errChan)
		}

		// error reporting
		go func(errChan <-chan error) {
			for {
				err := <-errChan
				log.Println(err)
			}
		}(errChan)

		wg.Wait()

		log.Println("Worker: sleeping blog aggregator worker")
		time.Sleep(time.Second * time.Duration(60))
	}
}

func (cfg *apiConfig) aggBlog(wg *sync.WaitGroup, feed database.Feed, ch chan<- error) {
	defer wg.Done()

	rss, err := getRssData(feed.Url)
	if err != nil {
		ch <- err
		return
	}

	for _, item := range rss.Channel.Item {
		postParams, err := createPostParams(item, feed.ID)
		if err != nil {
			ch <- err
			continue
		}
		_, err = cfg.DB.CreatePost(context.Background(), postParams)

		if err != nil {
			switch err.(type) {
			// ignore duplicate key errors
			case *pq.Error:
				if err.Error() == "pq: duplicate key value violates unique constraint \"posts_url_key\"" {
					continue
				} else {
					ch <- err
					continue
				}
			}
			switch err {
			// ignore duplicate errors
			// TODO: update the entry if I can't insert
			case sql.ErrNoRows:
			default:
				ch <- err
				continue
			}
		}
	}

	err = cfg.DB.MarkFeedFetch(context.Background(), database.MarkFeedFetchParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: feed.ID,
	})
	if err != nil {
		ch <- err
		return
	}
}

func createPostParams(item RssItem, feedID uuid.UUID) (database.CreatePostParams, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return database.CreatePostParams{}, err
	}
	pubDate, err := tryParseDate(item.PubDate)
	publishedAt := sql.NullTime{
		Time:  pubDate,
		Valid: true,
	}
	// failed to parse PubDate so insert null
	if err != nil {
		log.Println(err)
		publishedAt = sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		}
	}
	description := sql.NullString{
		String: item.Description,
		Valid:  true,
	}
	if item.Description == "" {
		description = sql.NullString{
			String: "",
			Valid:  false,
		}
	}
	params := database.CreatePostParams{
		ID:          id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
		Title:       item.Title,
		Url:         item.Link,
		Description: description,
		PublishedAt: publishedAt,
		FeedID:      feedID,
	}
	return params, nil
}

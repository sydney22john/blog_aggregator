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
)

// Rss xml structure
type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Atom    string   `xml:"atom,attr"`
	Version string   `xml:"version,attr"`
	Channel struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description   string `xml:"description"`
		Generator     string `xml:"generator"`
		Language      string `xml:"language"`
		LastBuildDate string `xml:"lastBuildDate"`
		Item          []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			PubDate     string `xml:"pubDate"`
			Guid        string `xml:"guid"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
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

	// TODO wg could potentially panic I believe so set something up to restart the worker function
	for {
		log.Println("Worker: getting top 10 feeds from db")
		feeds, err := cfg.DB.GetNextFeedsToFetch(context.Background(), int32(processFeeds))
		if err != nil {
			log.Println(err)
		}

		// TODO handle goroutine errors with an errors channel
		log.Println("Worker: starting go routines")
		for _, feed := range feeds {
			wg.Add(1)
			go cfg.aggBlog(&wg, feed)
		}

		wg.Wait()

		log.Println("Worker: sleeping blog aggregator worker")
		time.Sleep(time.Second * time.Duration(60))
	}
}

func (cfg *apiConfig) aggBlog(wg *sync.WaitGroup, feed database.Feed) error {
	defer wg.Done()

	log.Println(feed.Name)
	err := cfg.DB.MarkFeedFetch(context.Background(), database.MarkFeedFetchParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: feed.ID,
	})
	if err != nil {
		return err
	}
	return nil
}

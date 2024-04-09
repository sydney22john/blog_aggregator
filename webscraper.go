package main

import (
	"encoding/xml"
	"io"
	"net/http"
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

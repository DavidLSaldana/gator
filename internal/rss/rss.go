package rss

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"descrtiption"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	feed := RSSFeed{}
	rssFeedPtr := &feed

	//figuring this out. expecting errors here. Things worked before, so start
	//looking in all of this function below
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return rssFeedPtr, err
	}

	newClient := http.Client{}
	req.Header.Set(req.UserAgent(), "gator")

	resp, err := newClient.Do(req)
	if err != nil {
		return rssFeedPtr, err
	}

	reader, err := io.ReadAll(resp.Body)
	if err != nil {
		return rssFeedPtr, err
	}

	err = xml.Unmarshal(reader, rssFeedPtr)
	if err != nil {
		return rssFeedPtr, err
	}

	return rssFeedPtr, nil
}

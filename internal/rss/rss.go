package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
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

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	feed := RSSFeed{}
	rssFeedPtr := &feed

	//figuring this out. expecting errors here. Things worked before, so start
	//looking in all of this function below
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return rssFeedPtr, err
	}

	fmt.Println("This is happening - 1")
	client := http.Client{}
	req.Header.Set(req.UserAgent(), "gator")
	fmt.Println("This is happening - 2")

	fmt.Println("This is happening - 3")
	//error on do request, invalid header ""
	resp, err := client.Do(req)
	if err != nil {
		return rssFeedPtr, err
	}
	fmt.Println("This is happening - 4")

	reader, err := io.ReadAll(resp.Body)
	if err != nil {
		return rssFeedPtr, err
	}

	err = xml.Unmarshal(reader, rssFeedPtr)
	if err != nil {
		return rssFeedPtr, err
	}

	return cleanup(rssFeedPtr), nil
}

// decode escaped HTML from entire feed
func cleanup(rssFeed *RSSFeed) *RSSFeed {
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssFeed.Channel.Item[i].Description)
	}
	return rssFeed
}

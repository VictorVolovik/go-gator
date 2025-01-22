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
		Description string    `xml:"description"`
		Items       []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("unable to create request, %w", err)
	}
	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("unable to fetch feed data, %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("unable to read response data, %w", err)
	}

	var rssFeed RSSFeed

	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("unable to parse response data, %w", err)
	}

	unescapeFeed(&rssFeed)

	return &rssFeed, nil
}

func unescapeFeed(f *RSSFeed) {
	f.Channel.Title = html.UnescapeString(f.Channel.Title)
	f.Channel.Description = html.UnescapeString(f.Channel.Description)

	for i, item := range f.Channel.Items {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		f.Channel.Items[i] = item
	}
}

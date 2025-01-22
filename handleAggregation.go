package main

import (
	"context"
	"fmt"

	"VictorVolovik/go-gator/internal/rss"
)

const exampleFeedUlr = "https://www.wagslane.dev/index.xml"

func handleAggregation(s *State, cmd Command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	feed, err := rss.FetchFeed(context.Background(), exampleFeedUlr)
	if err != nil {
		return err
	}

	fmt.Printf("feed: %v", feed)

	return nil
}

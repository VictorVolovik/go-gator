package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"VictorVolovik/go-gator/internal/rss"
)

func handleAggregation(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("incorrect time format, %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenReqs)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ticker := time.NewTicker(timeBetweenReqs)
	defer ticker.Stop()

	if err := scrapeFeeds(s); err != nil {
		fmt.Printf("Error fetching feed: %v\n", err)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nTerminating aggregation process...")
			return nil
		case <-ticker.C:
			if err := scrapeFeeds(s); err != nil {
				fmt.Printf("Error fetching feed: %v\n", err)
			}
		}
	}
}

func scrapeFeeds(s *State) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get next feed to scrape, %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched, %w", err)
	}

	feedData, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed, %w", err)
	}

	fmt.Println("******")

	fmt.Printf("Feed title: %s\n", feedData.Channel.Title)
	fmt.Printf("Feed description: %s\n", feedData.Channel.Description)
	fmt.Println("------")

	for _, post := range feedData.Channel.Items {
		fmt.Printf("Post title: %s\n", post.Title)
		fmt.Println("------")
	}

	fmt.Printf("Feed %s collected, %d posts found\n", nextFeed.Name, len(feedData.Channel.Items))

	fmt.Println("******")
	fmt.Println("")

	return nil
}

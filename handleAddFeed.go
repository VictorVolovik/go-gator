package main

import (
	"context"
	"fmt"

	"VictorVolovik/go-gator/internal/database"
)

func handleAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	feedName := cmd.args[0]
	feedUrl := cmd.args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   feedName,
		Url:    feedUrl,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to create feed, %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to follow feed, %w", err)
	}

	fmt.Printf("Feed \"%s\" successfully added with url: %s\n", feed.Name, feed.Url)

	return nil
}

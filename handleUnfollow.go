package main

import (
	"context"
	"fmt"

	"VictorVolovik/go-gator/internal/database"
)

func handleUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	feedUrl := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("no feed found with such url, %w", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to follow feed, %w", err)
	}

	fmt.Printf("Feed \"%s\" successfully unfollowed\n", feed.Name)

	return nil
}

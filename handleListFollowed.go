package main

import (
	"context"
	"fmt"

	"VictorVolovik/go-gator/internal/database"
)

func handleListFollowed(s *State, cmd Command, user database.User) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("unable to get a list of user's followed feeds, %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No followed feeds found")
		return nil
	}

	fmt.Printf("Feeds followed by %s:\n", user.Name)
	fmt.Println("------")

	for _, feed := range feedFollows {
		fmt.Println(feed.FeedName)
		fmt.Println("------")
	}

	return nil
}

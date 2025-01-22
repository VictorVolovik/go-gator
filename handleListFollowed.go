package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func handleListFollowed(s *State, cmd Command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("current user not found")
		}
		return fmt.Errorf("unable to get current user info, %w", err)
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

	for _, feed := range feedFollows {
		fmt.Println("------")
		fmt.Println(feed.FeedName)
	}

	fmt.Println("------")

	return nil
}

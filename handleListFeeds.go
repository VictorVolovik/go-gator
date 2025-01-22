package main

import (
	"context"
	"fmt"
)

func handleListFeeds(s *State, cmd Command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get a list of all feeds, %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds added")
		return nil
	}

	fmt.Println("Feeds:")

	for _, feed := range feeds {
		fmt.Println("------")
		fmt.Printf("Feed \"%s\" - %s\n", feed.Name, feed.Url)
		fmt.Printf("Added by %s\n", feed.Username)
	}

	fmt.Println("------")

	return nil
}

package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"VictorVolovik/go-gator/internal/database"

	"github.com/google/uuid"
)

func handleAddFeed(s *State, cmd Command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("current user not found")
		}
		return fmt.Errorf("unable to get current user info, %w", err)
	}

	feedName := cmd.args[0]
	feedUrl := cmd.args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to create feed, %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to follow feed, %w", err)
	}

	fmt.Printf("Feed \"%s\" successfully added with url: %s\n", feed.Name, feed.Url)

	return nil
}

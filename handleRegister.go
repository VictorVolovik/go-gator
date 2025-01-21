package main

import (
	"VictorVolovik/go-gator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handleRegister(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.name)
	}

	username := cmd.args[0]

	_, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("unable to create user, %w", err)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("unable to handle login for a new user, %w", err)
	}

	fmt.Printf("user successfully created: %s\n", s.cfg.CurrentUserName)

	return nil
}

package main

import (
	"context"
	"fmt"
)

func handleRegister(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.name)
	}

	username := cmd.args[0]

	_, err := s.db.CreateUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("unable to create user, %w", err)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("unable to handle login for a new user, %w", err)
	}

	fmt.Printf("User successfully created: %s\n", s.cfg.CurrentUserName)

	return nil
}

package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func handleLogin(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.name)
	}

	username := cmd.args[0]

	user, err := s.db.GetUserByName(context.Background(), username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("unable to get user info, %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("unable to handle login with username provided, %w", err)
	}

	fmt.Printf("User successfully swithed: %s\n", s.cfg.CurrentUserName)

	return nil
}

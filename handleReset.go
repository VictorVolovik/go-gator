package main

import (
	"context"
	"fmt"
)

func handleReset(s *State, cmd Command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset, %w", err)
	}

	fmt.Println("App successfully reset")

	return nil
}

package main

import (
	"context"
	"fmt"
)

func handleListUsers(s *State, cmd Command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get a list of users, %w", err)
	}

	if len(users) == 0 {
		fmt.Println("No users available")
		return nil
	}

	fmt.Println("Users:")

	for _, user := range users {
		str := fmt.Sprintf("* %s", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			str += " (current)"
		}
		fmt.Println(str)
	}

	return nil
}

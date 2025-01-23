package main

import (
	"VictorVolovik/go-gator/internal/database"
	"context"
	"fmt"
	"strconv"
)

func handleBrowse(s *State, cmd Command, user database.User) error {
	limit := 2

	if len(cmd.args) == 1 {
		inputLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("incorrect limit")
		}
		limit = inputLimit
	} else if len(cmd.args) > 1 {
		return fmt.Errorf("usage: %s <limit - optional>", cmd.name)
	}

	posts, err := s.db.GetPosts(context.Background(), database.GetPostsParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("failed to get posts, %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found")
		return nil
	}

	fmt.Println("------")
	for _, post := range posts {
		fmt.Printf("Post title: %s\n", post.Title)
		fmt.Printf("Post description: %s\n", post.Description)
		fmt.Printf("Read more: %s\n", post.Url)
		fmt.Println("------")
	}

	return nil
}

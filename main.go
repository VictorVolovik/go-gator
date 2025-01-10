package main

import (
	"fmt"

	"VictorVolovik/go-gator/internal/config"
)

const username = "VictorVolovik"

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	err = cfg.SetUser(username)
	if err != nil {
		fmt.Println(err)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Gator config successfully set:")
	fmt.Printf("database url: %s\n", cfg.DbURL)
	fmt.Printf("username: %s\n", cfg.CurrentUserName)
}

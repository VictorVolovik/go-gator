package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"VictorVolovik/go-gator/internal/config"
	"VictorVolovik/go-gator/internal/database"
)

type State struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal("unable to open database connection: %w", err)
	}
	dbQueries := database.New(db)

	appState := State{
		cfg: &cfg,
		db:  dbQueries,
	}

	commands := Commands{
		registeredCommands: make(map[string]func(*State, Command) error),
	}

	commands.register("login", handleLogin)
	commands.register("register", handleRegister)
	commands.register("reset", handleReset)
	commands.register("users", handleListUsers)
	commands.register("agg", handleAggregation)
	commands.register("addfeed", handleAddFeed)

	args := os.Args

	if len(args) < 2 {
		log.Fatal("usage: cli <command> [args...]")
		return
	}

	commandName := args[1]
	commandArgs := args[2:]

	command := Command{
		name: commandName,
		args: commandArgs,
	}

	err = commands.run(&appState, command)
	if err != nil {
		log.Fatal(err)
	}
}

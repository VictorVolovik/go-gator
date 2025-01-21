package main

import "fmt"

type Command struct {
	name string
	args []string
}

type Commands struct {
	registeredCommands map[string]func(*State, Command) error
}

// Registers a new handler function for a command name
func (c *Commands) register(name string, f func(*State, Command) error) {
	c.registeredCommands[name] = f
}

// Runs a given command with the provided state if it exists
func (c *Commands) run(s *State, cmd Command) error {
	f, ok := c.registeredCommands[cmd.name]
	if !ok {
		return fmt.Errorf("command not found")
	}

	return f(s, cmd)
}

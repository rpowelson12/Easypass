package main

import (
	"errors"
)

type command struct {
	Name string
	Args []string
}

type commandEntry struct {
	handler     func(*state, command) error
	description string
}

type commands struct {
	registeredCommands map[string]commandEntry
}

func (c *commands) register(name string, description string, f func(*state, command) error) {
	c.registeredCommands[name] = commandEntry{
		handler:     f,
		description: description,
	}
}

func (c *commands) run(s *state, cmd command) error {
	entry, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("Command not found")
	}

	return entry.handler(s, cmd)
}

package main

import "fmt"

func handlerHelp(cmds *commands, s *state, cmd command) error {
	fmt.Println("Available commands:")
	for name, entry := range cmds.registeredCommands {
		fmt.Printf("  %-12s - %s\n", name, entry.description)
	}
	return nil
}

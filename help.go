package main

import "fmt"

func handlerHelp(s *state, cmd command) error {
	fmt.Printf(`
	help: displays a help message
	register: registers a new user
	login: logs in an existing user. Needs password
	users: lists all registered users
	generate: generates a new password for provided platform
	get: gets password for provided platform
	platforms: lists all platforms which have a password saved
	delete: delets the provided platform
	deactivate: deactivates provided user
	update: generates a new password for given platform
	upgrade: upgrade to the latest version
	`)
	return nil
}

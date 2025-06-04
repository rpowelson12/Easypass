package main

import (
	"errors"
	"fmt"
)

func commandCreate(args ...string) error {
	if len(args) != 2 {
		return errors.New("You must provide a username and password to create a user")
	}

	username := args[0]
	password := args[1]

	fmt.Printf("Creating user with username: %s and password: %s\n", username, password)

	return nil
}

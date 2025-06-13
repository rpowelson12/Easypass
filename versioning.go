package main

import (
	"fmt"

	"github.com/rpowelson12/Easypass/internal/version"
)

func handlerVersion(s *state, cmd command) error {
	fmt.Printf("Easypass version: %s\n", version.Version)
	fmt.Printf("Commit: %s\n", version.Commit)
	fmt.Printf("Build date: %s\n", version.Date)
	return nil
}

package main

import (
	"fmt"
	"github.com/rpowelson12/Easypass/internal/version"
)

func handlerVersion(s *state, cmd command) error {
	fmt.Printf("Easypass version: %s\n", version.Version)
	return nil
}

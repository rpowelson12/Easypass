package main

import (
	"fmt"
	"os"
)

func commandExit(args ...string) error {
	fmt.Println("Closing Easypass...")
	os.Exit(0)
	return nil
}

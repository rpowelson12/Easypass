package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/rpowelson12/Easypass/internal/database"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "", fmt.Errorf("cannot encrypt password: %v", err)
	}

	return string(hashedPassword), nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}
	name := cmd.Args[0]
	fmt.Print("Enter password:\n")
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("\nError reading password: %v\n", err)

	}

	hashedPassword, err := HashPassword(string(password))
	if err != nil {
		return fmt.Errorf("cannot hash password: %v", err)
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Password:  string(hashedPassword),
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully")
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]
	fmt.Print("Enter password: ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("\nError reading password: %v\n", err)

	}

	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}
	result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if result != nil {
		return errors.New("Incorrect user or password")
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully")
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not list users: %v", err)
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %v\n", user.Name)
	}
	return nil
}

func handlerDeleteUser(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <user name>", cmd.Name)
	}

	userName := cmd.Args[0]
	err := s.db.DeleteUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("cannot delete user: %v", err)
	}

	fmt.Printf("Successfully deleted %s", userName)

	return nil
}

func handlerUpdate(s *state, cmd command) error {

	fmt.Println("Upgrading Easypass...")

	modulePath := "github.com/rpowelson12/Easypass@latest"

	c := exec.Command("go", "install", modulePath)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		return fmt.Errorf("failed to upgrade Easypass: %w", err)
	}

	fmt.Println("Easypass upgraded!")
	return nil
}

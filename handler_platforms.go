package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"
	"github.com/rpowelson12/Easypass/internal/database"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@%^*+="

func generatePassword(length int) (string, error) {
	password := make([]byte, length)
	for i := range password {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[index.Int64()]
	}
	return string(password), nil
}

func handlerGenerate(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <platform>", cmd.Name)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("cannot find user: %v", err)
	}

	password, err := generatePassword(16)
	if err != nil {
		return fmt.Errorf("cannot generate password: %v", err)
	}
	user_id := user.ID
	platform := cmd.Args[0]

	platformObj, err := s.db.GeneratePassword(context.Background(), database.GeneratePasswordParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Platform:  platform,
		Password:  password,
		UserID:    user_id,
	})
	if err != nil {
		return fmt.Errorf("cannot generate platform object: %v", err)
	}
	clipboard.WriteAll(platformObj.Password)
	fmt.Println("Password copied to clipboard")

	return nil
}

func handlerGetPassword(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <platform>", cmd.Name)
	}
	platform := cmd.Args[0]
	user_id, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("cannot get user id: %v", err)
	}

	password, err := s.db.GetPassword(context.Background(), database.GetPasswordParams{
		Platform: platform,
		UserID:   user_id.ID,
	})

	fmt.Println("Password copied to clipboard!")
	clipboard.WriteAll(password)
	return nil
}

func handlerGetPlatforms(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}

	user_id, err := s.db.GetUserId(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("cannot get user id: %v", err)
	}

	platforms, err := s.db.GetPlatforms(context.Background(), user_id)

	for _, platform := range platforms {
		fmt.Println(platform.Platform)
	}

	return nil
}

func handlerDeletePlatform(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <platform>", cmd.Name)
	}

	platform := cmd.Args[0]

	err := s.db.DeletePlatform(context.Background(), platform)
	if err != nil {
		return fmt.Errorf("could not delete platform")
	}
	fmt.Printf("%s deleted successfully", platform)

	return nil
}

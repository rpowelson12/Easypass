package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/rpowelson12/Easypass/internal/database"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+="

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

	fmt.Println(password)

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

	fmt.Println(platformObj.Password)

	return nil
}

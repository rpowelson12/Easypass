package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
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

	//err = s.cfg.SetUser(user.Name)
	s.cfg.SetUser(user.Name)
	/*
		if err != nil {
			return fmt.Errorf("couldn't set current user: %w", err)
		}
	*/

	fmt.Println("User created successfully")
	checkVersions()
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
	s.cfg.SetUser(name)
	/*
		if err != nil {
			return fmt.Errorf("couldn't set current user: %w", err)
		}
	*/

	fmt.Println("User switched successfully")
	checkVersions()
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
	checkVersions()
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
	fmt.Println("🔄 Updating Easypass from latest GitHub release...")

	// Determine OS and Arch for the binary name
	osName := runtime.GOOS
	arch := runtime.GOARCH

	// Normalize arch to match your binary naming
	if arch == "amd64" {
		arch = "amd64"
	} else if arch == "arm64" || arch == "aarch64" {
		arch = "arm64"
	} else {
		return fmt.Errorf("unsupported architecture: %s", arch)
	}

	binaryName := fmt.Sprintf("easypass_%s_%s", osName, arch)

	// GitHub latest release download URL
	url := fmt.Sprintf("https://github.com/rpowelson12/Easypass/releases/latest/download/%s", binaryName)

	fmt.Printf("📦 Downloading %s...\n", url)

	// Download binary to temp file
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download binary: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to download binary, status: %s", resp.Status)
	}

	tmpFile, err := os.CreateTemp("", "easypass-update-*")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to temp file: %w", err)
	}

	if err := tmpFile.Chmod(0755); err != nil {
		return fmt.Errorf("failed to set executable permission: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	// Find current executable path
	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to find current executable: %w", err)
	}

	// Backup current executable (optional)
	backupPath := execPath + ".bak"
	err = copyFile(execPath, backupPath)
	if err != nil {
		return fmt.Errorf("failed to backup current executable: %w", err)
	}

	// Replace current executable with new binary
	err = os.Rename(tmpFile.Name(), execPath)
	if err != nil {
		// Might fail due to permissions, try to copy manually
		err = copyFile(tmpFile.Name(), execPath)
		if err != nil {
			return fmt.Errorf("failed to replace executable: %w", err)
		}
	}

	fmt.Println("✅ Easypass upgraded successfully!")

	return nil
}

// copyFile copies src to dst (overwrites if exists)
func copyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dst, input, 0755)
	return err
}

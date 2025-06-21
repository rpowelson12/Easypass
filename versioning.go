package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rpowelson12/Easypass/internal/version"
)

func handlerVersion(s *state, cmd command) error {
	fmt.Printf("Easypass version: %s\n", version.Version)
	fmt.Printf("Commit: %s\n", version.Commit)
	fmt.Printf("Build date: %s\n", version.Date)

	checkVersions()
	return nil
}

func checkVersions() error {
	latest, err := getLatestVersion()
	if err != nil {
		return fmt.Errorf("could not get latest version: %v", err)
	}
	if version.Version != latest {
		fmt.Println("New version available! <easypass update> to get new version")
	}
	return nil
}

func getLatestVersion() (string, error) {
	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get("https://api.github.com/repos/rpowelson12/Easypass/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}
	return data.TagName, nil
}

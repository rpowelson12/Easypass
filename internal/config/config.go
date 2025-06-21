package config

import (
	//"encoding/json"
	//"os"
	//"path/filepath"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

const configFileName = ".easypassconfig.json"

type Config struct {
	KEY             string `env:"ENCRYPTION_KEY"`
	DBURL           string `env:"DB_URL"`
	CurrentUserName string `env:"CURRENT_USER"`
}

func Load() (*Config, error) {

	_ = godotenv.Load()

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (cfg *Config) SetUser(userName string) {
	cfg.CurrentUserName = userName
}

/*func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
*/

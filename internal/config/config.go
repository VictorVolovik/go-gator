package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	configFileName        = ".gatorconfig.json"
	configFilePermissions = 0644 // -rw-r--r--
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username

	err := write(*cfg)
	if err != nil {
		return fmt.Errorf("error setting user's name to config, %w", err)
	}

	return nil
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error getting config file path, %w", err)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err = json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("error parsing config file: %w", err)
	}

	return cfg, nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path, %w", err)
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error converting config file to json: %w", err)
	}

	err = os.WriteFile(filePath, data, configFilePermissions)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting user's home directory: %w", err)
	}

	configFilePath := homeDir + "/" + configFileName

	return configFilePath, nil
}

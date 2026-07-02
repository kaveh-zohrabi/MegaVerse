package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	APIKey  string `json:"api_key"`
	BaseURL string `json:"base_url"`
}

const configDir = ".megaverse"
const configFile = "config.json"

func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configDir, configFile), nil
}

func Load() (*Config, error) {
	path, err := GetConfigPath()
	if err != nil {
		return &Config{BaseURL: "http://localhost:8080"}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return &Config{BaseURL: "http://localhost:8080"}, nil
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return &Config{BaseURL: "http://localhost:8080"}, nil
	}

	return &cfg, nil
}

func Save(cfg *Config) error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

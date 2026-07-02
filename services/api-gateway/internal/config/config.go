package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port           string
	AuthServiceURL string
	UserServiceURL string
	RateLimit      int
	RateWindowSec  int
}

func Load() *Config {
	return &Config{
		Port:           getEnv("PORT", "8080"),
		AuthServiceURL: getEnv("AUTH_SERVICE_URL", "http://localhost:8081"),
		UserServiceURL: getEnv("USER_SERVICE_URL", "http://localhost:8082"),
		RateLimit:      getEnvInt("RATE_LIMIT", 1000),
		RateWindowSec:  getEnvInt("RATE_WINDOW", 60),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

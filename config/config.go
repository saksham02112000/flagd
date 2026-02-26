package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	RedisURL    string
	Port        int
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := &Config{}

	required := []string{"DATABASE_URL", "REDIS_URL", "PORT"}
	for _, key := range required {
		if os.Getenv(key) == "" {
			return nil, fmt.Errorf("required env var %q is not set", key)
		}
	}

	cfg.DatabaseURL = os.Getenv("DATABASE_URL")
	cfg.RedisURL = os.Getenv("REDIS_URL")

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("PORT must be a valid integer: %w", err)
	}
	cfg.Port = port

	return cfg, nil
}

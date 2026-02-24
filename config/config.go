package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	RedisURL    string
	Port        string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		RedisURL:    os.Getenv("REDIS_URL"),
		Port:        os.Getenv("PORT"),
	}

	required := map[string]*string{
		"DATABASE_URL": &cfg.DatabaseURL,
		"REDIS_URL":    &cfg.RedisURL,
		"PORT":         &cfg.Port,
	}
	for key, dest := range required {
		val := os.Getenv(key)
		if val == "" {
			return nil, fmt.Errorf("required env var %q is not set", key)
		}
		*dest = val
	}

	return cfg, nil
}

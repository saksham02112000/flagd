package main

import (
	"flagd/config"
	"flagd/config/client"
	"flagd/internal/logger"
	"flagd/internal/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	appLogger := logger.New(os.Getenv("ENV"))

	postgres := client.ConnectPostgres(*cfg)
	defer postgres.Close()

	redis := client.ConnectRedis(*cfg)
	defer redis.Close()

	c := routes.Build(cfg, postgres, redis, appLogger)
	routes.SetupHandlers(app, c)

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Fatal(app.Listen(addr))
}

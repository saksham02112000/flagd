package main

import (
	"flagd/config"
	"flagd/config/client"
	"flagd/internal/logger"
	"flagd/internal/routes"
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
	redis := client.ConnectRedis(*cfg)

	c := routes.Build(cfg, postgres, redis, appLogger)
	routes.SetupHandlers(app, c)

	app.Listen(":8080")

}

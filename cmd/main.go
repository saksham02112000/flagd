package main

import (
	"flagd/config"
	"flagd/config/client"
	"flagd/internal/routes"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	postgres := client.ConnectPostgres(*cfg)
	redis := client.ConnectRedis(*cfg)

	c := routes.Build(cfg, postgres, redis)
	routes.SetupHandlers(app, c)

	app.Listen(":8080")

}

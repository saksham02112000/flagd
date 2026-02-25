package routes

import (
	"flagd/internal/logger/middleware"

	"github.com/gofiber/fiber/v3"
)

func SetupHandlers(app *fiber.App, c *Container) {
	app.Use(middleware.RequestLogger(c.Logger))

	api := app.Group("/api/v1")

	api.Get("/flag/:id", c.FlagHandler.GetById)
	api.Post("/flag", c.FlagHandler.Create)
}

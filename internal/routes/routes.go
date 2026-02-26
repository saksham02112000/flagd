package routes

import (
	"flagd/internal/logger/middleware"

	"github.com/gofiber/fiber/v3"
)

func SetupHandlers(app *fiber.App, c *Container) {
	app.Use(middleware.RequestLogger(c.Logger))

	api := app.Group("/api/v1")
	flags := api.Group("/flags")

	flags.Get("/", c.FlagHandler.GetAll)
	flags.Get("/:id", c.FlagHandler.GetById)
	flags.Post("/", c.FlagHandler.Create)
	flags.Delete("/:id", c.FlagHandler.Delete)
	flags.Patch("/:id/environments/:envSlug", c.FlagHandler.Toggle)
}

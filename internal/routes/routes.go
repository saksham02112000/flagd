package routes

import (
	"github.com/gofiber/fiber/v3"
)


func SetupHandlers(app *fiber.App, c *Container){

	api := app.Group("/api/v1")

	api.Get("/flag/:id", c.FlagHandler.GetById)


	
}
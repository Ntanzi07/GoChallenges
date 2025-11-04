package routes

import (
	"Ex4/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/tasks")

	api.Get("/", handlers.GetAll)
	api.Get("/:id", handlers.GetById)
	api.Post("/", handlers.Create)
	api.Put("/:id", handlers.Update)
	api.Delete("/:id", handlers.Delete)
}

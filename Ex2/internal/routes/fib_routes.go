package routes

import (
	"Ex2/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/all", handlers.GetAll)
	app.Get("/fib/:n", handlers.GetFibNum)
}

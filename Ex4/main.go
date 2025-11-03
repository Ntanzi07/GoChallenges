package main

import (
	"Ex4/data"
	"Ex4/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	err := data.LoadTasks()
	if err != nil {
		log.Fatalf("Erro ao carregar tarefas: %v", err)
	}

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}

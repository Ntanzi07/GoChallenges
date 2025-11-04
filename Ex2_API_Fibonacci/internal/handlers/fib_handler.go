package handlers

import (
	"Ex2/internal/data"
	"Ex2/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	results := data.GetAll()

	return c.JSON(fiber.Map{
		"queries": results,
	})
}

func GetFibNum(c *fiber.Ctx) error {
	numStr := c.Params("n")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "num inválido"})
	}
	if num < 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Numero negativo é invalido"})
	}

	result, ok := services.CalculaFib(num)

	return c.JSON(fiber.Map{
		"done": ok,
		"fib":  result,
	})
}

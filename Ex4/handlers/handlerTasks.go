package handlers

import (
	"Ex4/data"
	"Ex4/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {

	query, err := strconv.ParseBool(c.Query("completed"))
	if err != nil {
		tasks := data.GetAll()
		return c.JSON(tasks)
	}

	filtered := data.GetByCompleted(query)

	return c.JSON(filtered)
}

func GetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "task not found"})
	}

	task, err := data.GetById(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "task not found"})
	}

	return c.JSON(task)
}

func Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	var update models.Task
	if err := c.BodyParser(&update); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
	}

	update.ID = id
	task, err := data.Update(update)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "task not found"})
	}

	return c.JSON(task)
}

func Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := data.Delete(id); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "task not found"})
	}

	return c.JSON(fiber.Map{"message": "Tarefa removida com sucesso"})
}

func Create(c *fiber.Ctx) error {

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
	}
	if input.Title == "" {
		return c.Status(400).JSON(fiber.Map{"error": "title is required"})
	}
	if input.Description == "" {
		return c.Status(400).JSON(fiber.Map{"error": "description is required"})
	}

	newTask, err := data.Create(input.Title, input.Description)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "task not created"})
	}

	return c.Status(201).JSON(newTask)
}

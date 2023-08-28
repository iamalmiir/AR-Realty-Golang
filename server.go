package main

import (
	"golabs/db"
	"golabs/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	db := db.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		var allUsers []models.User
		db.Find(&allUsers)

		return c.JSON(allUsers)
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User
		err := c.BodyParser(&user)
		if err != nil {
			return c.Status(400).JSON(err)
		}
		user.ID = uuid.New()
		user_exists := db.Create(&user)

		if user_exists.Error != nil {
			return c.Status(400).JSON("User already exists")
		}

		return c.Status(200).JSON(user)
	})

	app.Listen(":8080")
}

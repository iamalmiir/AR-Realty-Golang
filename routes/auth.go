package routes

import (
	"golabs/db"
	"golabs/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user models.User
	err := c.BodyParser(&user)
	db := db.DB
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	userStorage := models.UserStorage{Conn: db}

	err = userStorage.NewUser(&user)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.Status(200).SendString("User created!")
}

func GetAllUsers(c *fiber.Ctx) error {
	db := db.DB
	userStorage := models.UserStorage{Conn: db}

	users, err := userStorage.GetUsers()
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.Status(200).JSON(users)
}

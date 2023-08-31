package routes

import (
	"golabs/db"
	"golabs/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Register(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user models.User
	err := c.BodyParser(&user)

	db := db.DB
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	userStorage := models.UserStorage{Conn: db}
	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create user",
			"data":    err.Error(),
		})
	}
	user.Password = hash
	err = userStorage.NewUser(&user)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.Status(200).SendString("User created!")
}

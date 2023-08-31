package routes

import (
	"golabs/db"
	"golabs/models"
	"golabs/utils"

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

	if err != nil {
		return c.Status(400).JSON(utils.ServerResponse(400, "Something is wrong with your request. Please check your request and try again."))
	}

	db := db.DB
	userStorage := models.UserStorage{Conn: db}
	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(utils.ServerResponse(500, "Something went wrong. Please try again later or contact us."))
	}
	user.Password = hash
	err = userStorage.NewUser(&user)
	if err != nil {
		return c.Status(500).JSON(utils.ServerResponse(500, "Something went wrong. Please check your data or try again later."))
	}

	return c.Status(200).JSON(utils.ServerResponse(200, "User registered successfully", user))
}

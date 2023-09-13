package services

import (
	"golabs/db"
	"golabs/models"
	"golabs/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

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

func GetUserByEmail(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user models.User
	err := c.BodyParser(&user)
	sessionID := c.Cookies("session_id")
	println(sessionID)

	if err != nil {
		return c.Status(400).JSON(utils.ServerResponse(400, "Something is wrong with your request. Please check your request and try again."))
	}

	db := db.DB
	userStorage := models.UserStorage{Conn: db}
	retrievedUser, err := userStorage.GetUserByEmail(user.Email)
	if err != nil {
		return c.Status(500).JSON(utils.ServerResponse(500, "Something went wrong. Please try again later or contact us."))
	}

	// Now, assign the retrieved user to your 'user' variable.
	user = *retrievedUser

	return c.Status(200).JSON(utils.ServerResponse(200, "User retrieved successfully", user))
}

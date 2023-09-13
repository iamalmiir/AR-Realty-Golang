package services

import (
	"context"
	"encoding/json"
	"fmt"
	"golabs/db"
	"golabs/models"
	"golabs/utils"
	"net/mail"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserData struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

func Login(c *fiber.Ctx) error {
	input := new(LoginInput)
	var userData UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ServerResponse(400, "Error on login request", err))
	}

	identity := input.Email
	pass := input.Password
	userModel := &models.User{}

	db, rdb := db.DB, db.Rdb
	userStorage := models.UserStorage{Conn: db}
	var err error
	if isEmail(identity) {
		userModel, err = userStorage.GetUserByEmail(identity)
	}

	if userModel == nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ServerResponse(404, "User not found", err))
	} else {
		userData = UserData{
			ID:        userModel.ID,
			FirstName: userModel.FirstName,
			LastName:  userModel.LastName,
			Email:     userModel.Email,
			Password:  userModel.Password,
		}
	}

	if !CheckPasswordHash(pass, userData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ServerResponse(401, "Invalid password", err))
	}

	sessionID := uuid.NewString()
	jsonData, err := json.Marshal(userData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerResponse(500, "Internal server error!", err))
	}
	err = rdb.Set(context.Background(), sessionID, jsonData, 24*time.Hour).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerResponse(500, "Internal server error!!", err))
	}

	c.Response().Header.Set("Authorization", fmt.Sprintf("Bearer %s", sessionID))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":    map[string]interface{}{"first_name": userData.FirstName, "last_name": userData.LastName, "email": userData.Email},
		"status":  "success",
		"message": "Login success",
	})
}

func GetSessionData(c *fiber.Ctx) error {
	sessionID := c.Get("Authorization")

	if sessionID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ServerResponse(400, "Session not found", nil))
	}

	rdb := db.Rdb
	println(sessionID[7:])
	sessionData, err := rdb.Get(context.Background(), sessionID[7:]).Result()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerResponse(500, "Internal server error--", err))
	}

	var userData UserData
	err = json.Unmarshal([]byte(sessionData), &userData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerResponse(500, "Internal server error", err))
	}

	return c.Status(fiber.StatusOK).JSON(userData)
}

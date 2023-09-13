package services

import (
	"fmt"
	"golabs/middleware"
	"golabs/models"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Storage        *models.UserStorage
	SessionManager *middleware.SessionManager
}

func NewUserHandler(storage *models.UserStorage, sessionManager *middleware.SessionManager) *UserHandler {
	return &UserHandler{
		Storage:        storage,
		SessionManager: sessionManager,
	}
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

type signInRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (u *UserHandler) SignInUser(c *fiber.Ctx) error {
	var user signInRequestBody

	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	fmt.Println(user)

	// validate the user struct
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		return err
	}

	// sign the user in
	sessionId, err := u.SessionManager.SignIn(user.Email, user.Password)
	if err != nil {
		return err
	}

	// set the session id as a header
	c.Response().Header.Set("Authorization", fmt.Sprintf("Bearer %s", sessionId))

	return c.JSON(fiber.Map{"success": true})
}

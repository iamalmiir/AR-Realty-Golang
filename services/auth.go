package services

import (
	"golabs/config"
	"golabs/db"
	"golabs/models"
	"golabs/utils"
	"net/mail"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

func Login(c *fiber.Ctx) error {
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

	input := new(LoginInput)
	var userData UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ServerResponse(400, "Error on login request", err))
	}

	identity := input.Email
	pass := input.Password
	userModel := &models.User{}

	db := db.DB
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

	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userData.ID
	claims["email"] = userData.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(config.GetEnv("SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ServerResponse(500, "Internal server error", err))
	}
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HTTPOnly = true

	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(utils.ServerResponse(200, "Success", UserData{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
	}, t))
}

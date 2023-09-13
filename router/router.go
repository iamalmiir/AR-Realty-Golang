package router

import (
	"golabs/config"
	"golabs/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Auth
	auth := app.Group("/auth", logger.New())
	auth.Use(encryptcookie.New(encryptcookie.Config{
		Key: config.GetEnv("SECRET"),
	}))
	auth.Post("/login", services.Login)

	// User routes
	user := app.Group("/user", logger.New())
	user.Post("/register", services.Register)
	user.Post("/", services.GetUserByEmail)
	user.Get("/me", services.GetSessionData)
}

package main

import (
	"log"

	"golabs/db"
	"golabs/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db.ConnectDB()

	app := fiber.New()
	app.Use(logger.New())

	app.Post("/register", routes.Register)
	app.Get("/users", routes.GetAllUsers)

	log.Fatal(app.Listen(":3000"))
}

package main

import (
	"log"

	"golabs/db"
	"golabs/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	// Initialize your database connection
	db.ConnectDB()

	// Set up your routes
	router.SetupRoutes(app)

	// Start the Fiber app
	log.Fatal(app.Listen(":8080"))
}

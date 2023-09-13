package main

import (
	"golabs/db"
	"golabs/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	// Initialize your database connection
	db.ConnectDB()
	db.ConnectRedis()

	// Set up your routes with the additional arguments
	router.SetupRoutes(app)

	// Start the Fiber app
	log.Fatal(app.Listen(":8080"))
}

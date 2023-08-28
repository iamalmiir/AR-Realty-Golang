package main

import (
	"golabs/models"
	"golabs/routes"
)

func main() {
	models.SetupDB()
	routes.Setup()
}

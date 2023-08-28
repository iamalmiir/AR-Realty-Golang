package db

import (
	"log"
	"os"

	"golabs/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading.env file, %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Error connecting to database: %s\n", err.Error())
		return nil
	}

	db.AutoMigrate(&models.User{})
	log.Println("Successfully connected to database")

	return db
}

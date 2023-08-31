package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	}

	DB, err = sqlx.Connect("mysql", os.Getenv("DSN"))
	if err != nil {
		panic("Error connecting to database")
	}

	log.Println("Connected to DB")
}

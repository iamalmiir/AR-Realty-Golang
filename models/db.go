package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func SetupDB() {
	dotErr := godotenv.Load()
	if dotErr != nil {
		panic("Error loading.env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	var err error
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to database")
}

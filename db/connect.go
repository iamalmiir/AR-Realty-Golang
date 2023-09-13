package db

import (
	"golabs/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB() {
	err := error(nil)
	DB, err = sqlx.Connect("mysql", config.GetEnv("DSN"))
	if err != nil {
		panic("Error connecting to database")
	}

	log.Println("Connected to DB")
}

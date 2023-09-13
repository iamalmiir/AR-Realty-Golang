package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		os.Exit(1)
	}

	return os.Getenv(key)
}

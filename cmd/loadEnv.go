package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, string, string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	return os.Getenv("USER_PG"), os.Getenv("PASSWORD_PG"), os.Getenv("DB_PG")
}

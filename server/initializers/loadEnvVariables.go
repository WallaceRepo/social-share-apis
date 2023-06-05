package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() (string, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB := os.Getenv("DB")
	APP_ID := os.Getenv("APP_ID")
	APP_SECRET := os.Getenv("APP_SECRET")
	REDIRECT_URL := os.Getenv("REDIRECT_URL")

	return DB, APP_ID, APP_SECRET, REDIRECT_URL
}


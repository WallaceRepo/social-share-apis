package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	
)

func LoadEnvVariables() (string, string, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB := os.Getenv("DB")
	CONSUMER_KEY :=os.Getenv("TWIT_CONSUMER_KEY")
	CONSUMER_SECRET := os.Getenv("TWIT_CONSUMER_SECRET")
	ACCESS_TOKEN := os.Getenv("TWIT_ACCESS_TOKEN")
	ACCESS_SECRET := os.Getenv("TWIT_TOKEN_SECRET")

	return DB, CONSUMER_KEY, CONSUMER_SECRET, ACCESS_TOKEN, ACCESS_SECRET
}


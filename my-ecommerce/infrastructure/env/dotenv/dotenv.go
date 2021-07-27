package dotenv

import (
	"github.com/joho/godotenv"
	"log"
)

func InitEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error to load .env file err: %v", err)
	}
}

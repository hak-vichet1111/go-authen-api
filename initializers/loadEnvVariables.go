package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	// envPath := "./.env"
	// envConfig, err := config.LoadEnvConfig(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

}
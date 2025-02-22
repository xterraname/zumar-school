package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return err
	}
	return nil
}

func GetPort() string {
	port, exists := os.LookupEnv("PORT")

	if !exists {
		port = "8080"
	}

	return port
}

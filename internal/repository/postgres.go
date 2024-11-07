package repository

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadDBConfig() (*DBConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	config := &DBConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
	return config, nil
}

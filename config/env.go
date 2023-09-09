package config

import (
	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
}

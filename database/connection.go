package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Instance *gorm.DB

func ConnectToDB() {
	dsn := os.Getenv("DATABASE_URL")

	var err error

	Instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
}

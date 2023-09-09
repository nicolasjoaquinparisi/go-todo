package database

import (
	"fmt"
	"go-todo/api/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Print("Failed to generate migration")

		return
	}
}

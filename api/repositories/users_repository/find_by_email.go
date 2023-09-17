package users_repository

import (
	"fmt"
	"go-todo/api/models"
	"go-todo/database"
)

func FindByEmail(email string) (*models.User, error) {
	fmt.Println("Invoked - Users repository FindByEmail")

	var user *models.User

	fmt.Println("Users Repository - Database Find First")
	result := database.Instance.First(&user, "email = ?", email)

	// return query error
	if result.Error != nil {
		fmt.Printf("Users Repository - %v\n", result.Error)
		return nil, result.Error
	}

	return user, nil
}

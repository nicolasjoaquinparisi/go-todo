package users_repository

import (
	"fmt"
	"github.com/google/uuid"
	"go-todo/api/models"
	"go-todo/database"
)

func FindById(id uuid.UUID) (*models.User, error) {
	fmt.Println("Invoked - Users repository FindById")

	var user *models.User

	result := database.Instance.First(&user, "id = ?", id)

	// return query error
	if result.Error != nil {
		fmt.Printf("Users Repository - %v\n", result.Error)
		return nil, result.Error
	}

	return user, nil
}

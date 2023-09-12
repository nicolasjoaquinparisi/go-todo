package users_repository

import (
	"go-todo/api/models"
	"go-todo/database"
)

func FindByEmail(email string) (*models.User, error) {
	var user *models.User

	result := database.Instance.First(&user, "email = ?", email)

	// return query error
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

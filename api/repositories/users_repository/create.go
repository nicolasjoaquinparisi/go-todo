package users_repository

import (
	"errors"
	"go-todo/api/models"
	"go-todo/database"
)

func Create(email string, password string, firstName string, lastName string) (*models.User, error) {
	user := models.User{Email: email, Password: password, FirstName: firstName, LastName: lastName}

	result := database.DB.Create(&user)

	if result.Error != nil {
		return nil, errors.New("failed to create user")
	}

	return &user, nil
}

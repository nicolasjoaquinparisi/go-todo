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

func FindByEmail(email string) (*models.User, error) {
	var user *models.User

	result := database.DB.First(&user, "email = ?", email)

	// return query error
	if result.Error != nil {
		return nil, result.Error
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

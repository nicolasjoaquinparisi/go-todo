package users_repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go-todo/api/models"
	"go-todo/database"
)

func Create(email string, password string, firstName string, lastName string) (*models.User, error) {
	fmt.Println("Invoked: Users repository Create")

	user := models.User{Email: email, Password: password, FirstName: firstName, LastName: lastName}

	user.ID = uuid.New()

	result := database.Instance.Create(&user)

	if result.Error != nil {
		return nil, errors.New("failed to create user")
	}

	return &user, nil
}

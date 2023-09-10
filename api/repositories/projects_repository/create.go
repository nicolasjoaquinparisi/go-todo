package projects_repository

import (
	"errors"
	"go-todo/api/models"
	"go-todo/database"
)

func Create(name string, description string, userId uint) (*models.Project, error) {
	project := models.Project{Name: name, Description: description, UserID: userId}

	result := database.DB.Create(&project)

	if result.Error != nil {
		return nil, errors.New("failed to create project")
	}

	return &project, nil
}

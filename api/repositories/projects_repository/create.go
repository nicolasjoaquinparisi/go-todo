package projects_repository

import (
	"errors"
	"github.com/google/uuid"
	"go-todo/api/models"
	"go-todo/database"
)

func Create(name string, description string, userId uuid.UUID) (*models.Project, error) {
	project := models.Project{Name: name, Description: description, UserID: userId}

	project.ID = uuid.New()

	result := database.Instance.Create(&project)

	if result.Error != nil {
		return nil, errors.New("failed to create project")
	}

	return &project, nil
}

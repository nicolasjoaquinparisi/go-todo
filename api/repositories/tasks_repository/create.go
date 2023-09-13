package tasks_repository

import (
	"errors"
	"github.com/google/uuid"
	"go-todo/api/models"
	"go-todo/database"
)

func Create(name string, description string, projectId uuid.UUID) (*models.Task, error) {
	task := models.Task{Name: name, Description: description, ProjectID: projectId}

	task.ID = uuid.New()

	result := database.Instance.Create(&task)

	if result.Error != nil {
		return nil, errors.New("failed to create task")
	}

	return &task, nil
}

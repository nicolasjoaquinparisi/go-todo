package tasks_repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go-todo/api/models"
	"go-todo/database"
)

func Create(name string, description string, projectId uuid.UUID) (*models.Task, error) {
	fmt.Println("Invoked: Tasks repository Create")

	task := models.Task{Name: name, Description: description, ProjectID: projectId}

	task.ID = uuid.New()

	result := database.Instance.Create(&task)

	if result.Error != nil {
		fmt.Println("Error creating task in the database:", result.Error)
		return nil, errors.New("failed to create task")
	}

	return &task, nil
}

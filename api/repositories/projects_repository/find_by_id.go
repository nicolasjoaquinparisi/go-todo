package projects_repository

import (
	"go-todo/api/models"
	"go-todo/database"
)

func FindById(id uint) (*models.Project, error) {
	var project *models.Project

	result := database.Instance.First(&project, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return project, nil
}

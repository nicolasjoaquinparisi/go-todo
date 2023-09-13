package projects_repository

import (
	"github.com/google/uuid"
	"go-todo/api/models"
	"go-todo/database"
)

func IsDuplicateProjectNameForUser(projectName string, userId uuid.UUID) (bool, error) {
	var project *models.Project

	result := database.Instance.First(&project, "AND name = ? AND user_id = ?", projectName, userId)

	if result.Error != nil {
		return false, result.Error
	}

	return project != nil, nil
}

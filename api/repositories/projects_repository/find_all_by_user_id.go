package projects_repository

import (
	"github.com/google/uuid"
	"go-todo/api/models"
	"go-todo/database"
)

func FindAllByUserId(userId uuid.UUID) ([]*models.Project, error) {
	var projects []*models.Project

	result := database.Instance.Find(&projects, "user_id = ?", userId)

	if result.Error != nil {
		return nil, result.Error
	}

	return projects, nil
}

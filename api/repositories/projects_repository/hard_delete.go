package projects_repository

import (
	"go-todo/api/models"
	"go-todo/database"
)

func HardDelete(project *models.Project) error {
	if err := database.Instance.Unscoped().Delete(project).Error; err != nil {
		return err
	}
	return nil
}

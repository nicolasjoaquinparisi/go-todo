package projects_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-todo/api/models"
	"go-todo/api/repositories/projects_repository"
	"go-todo/api/utils/responses/projects_responses"
	"net/http"
)

func FindById(c *gin.Context) {
	var result *models.Project

	projectId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Internal server error"})
		return
	}

	result, err = projects_repository.FindById(projectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Internal server error"})
		return
	}

	project := projects_responses.FindByIdResponse{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Tasks:       result.Tasks,
	}

	c.JSON(http.StatusOK, gin.H{"project": project})
}

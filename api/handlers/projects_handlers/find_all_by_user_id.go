package projects_handlers

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/models"
	"go-todo/api/repositories/projects_repository"
	"go-todo/api/utils/responses/projects_responses"
	"net/http"
)

func FindAllByUserId(c *gin.Context) {
	user, _ := c.Get("user")
	userId := user.(*models.User).ID

	result, err := projects_repository.FindAllByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Internal server error"})
		return
	}

	var projects []projects_responses.FindAllByUserIdResponse
	for _, project := range result {
		projectResponse := projects_responses.FindAllByUserIdResponse{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
		}
		projects = append(projects, projectResponse)
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

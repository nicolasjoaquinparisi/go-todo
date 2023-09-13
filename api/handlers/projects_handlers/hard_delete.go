package projects_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-todo/api/repositories/projects_repository"
	"net/http"
)

func HardDelete(c *gin.Context) {
	// get project id from request param and parse
	projectId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Internal server error"})
		return
	}

	// check if the project exists
	project, err := projects_repository.FindById(projectId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"description": "Project not found"})
		return
	}

	// hard delete project
	err = projects_repository.HardDelete(project)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"description": "Failed to delete project"})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"description": "Project deleted"})
}

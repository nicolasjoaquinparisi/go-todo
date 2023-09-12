package projects_handlers

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/repositories/projects_repository"
	"net/http"
	"strconv"
)

func HardDelete(c *gin.Context) {
	// get project id from request param and parse
	projectId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Internal server error"})
		return
	}

	// check if the project exists
	project, err := projects_repository.FindById(uint(projectId))
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

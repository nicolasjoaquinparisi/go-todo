package tasks_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-todo/api/repositories/tasks_repository"
	"net/http"
)

func Create(c *gin.Context) {
	var body struct {
		Name        string    `json:"name" validate:"required"`
		Description string    `json:"description"`
		ProjectId   uuid.UUID `json:"project_id" validate:"required"`
	}

	// map request body into body struct
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "Invalid payload"})
		return
	}

	// create project
	_, err := tasks_repository.Create(body.Name, body.Description, body.ProjectId)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"description": "Failed to create task"})
		return
	}

	// return response
	c.JSON(http.StatusCreated, gin.H{"description": "Task created"})
}

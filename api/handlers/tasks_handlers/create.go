package tasks_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-todo/api/repositories/projects_repository"
	"go-todo/api/repositories/tasks_repository"
	"go-todo/api/utils/responses/tasks_responses"
	"net/http"
)

func Create(c *gin.Context) {
	fmt.Println("Invoked: Tasks handler Create")

	var body struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
		ProjectId   string `json:"project_id" validate:"required"`
	}

	// map request body into body struct
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "Invalid payload"})
		return
	}

	// parse project id
	projectId, err := uuid.Parse(body.ProjectId)
	if err != nil {
		fmt.Println("Error parsing project_id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"description": "Invalid project_id"})
		return
	}

	// validate project
	project, err := projects_repository.FindById(projectId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"description": "Project not found"})
		return
	}

	// create task
	task, err := tasks_repository.Create(body.Name, body.Description, project.ID)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"description": "Failed to create task"})
		return
	}

	// format response
	taskResponse := tasks_responses.CreateHandlerResponse{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status,
		ProjectID:   task.ProjectID,
	}

	// return response
	c.JSON(http.StatusCreated, gin.H{"task": taskResponse})
}

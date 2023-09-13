package projects_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-todo/api/models"
	"go-todo/api/repositories/projects_repository"
	"net/http"
)

func Create(c *gin.Context) {
	var body struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}
	var validate *validator.Validate
	var err error

	user, _ := c.Get("user")
	userId := user.(models.User).ID

	// map request body into body struct
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "Invalid payload"})
		return
	}

	// validate body
	validate = validator.New()
	if err = validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate duplicates
	isProjectNameInUse, _ := projects_repository.IsDuplicateProjectNameForUser(body.Name, userId)
	if isProjectNameInUse == true {
		c.JSON(http.StatusConflict, gin.H{"description": "Project name already in use"})
		return
	}

	// create project
	_, err = projects_repository.Create(body.Name, body.Description, userId)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"description": "Failed to create project"})
		return
	}

	// return response
	c.JSON(http.StatusCreated, gin.H{"description": "Project created"})
}

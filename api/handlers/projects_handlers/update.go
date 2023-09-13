package projects_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-todo/api/models"
	"go-todo/api/repositories/projects_repository"
	"go-todo/api/utils/requests/projects_requests"
	"net/http"
)

func Update(c *gin.Context) {
	var body projects_requests.UpdateBodyStruct

	// get user id
	user, _ := c.Get("user")
	userId := user.(models.User).ID

	// get project id from request param and parse
	projectId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Internal server error"})
		return
	}

	// map request body into body struct
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "Invalid payload"})
		return
	}

	// check if the project exists
	project, err := projects_repository.FindById(projectId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"description": "Project not found"})
		return
	}

	// validate duplicates
	if body.Name != "" {
		isProjectNameInUse, _ := projects_repository.IsDuplicateProjectNameForUser(body.Name, userId)
		if isProjectNameInUse == true {
			c.JSON(http.StatusConflict, gin.H{"description": "Project name already in use"})
			return
		}
	}

	// update project
	_, err = projects_repository.Update(project, body)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"description": "Failed to update project"})
		return
	}

	// return response
	c.JSON(http.StatusCreated, gin.H{"description": "Project updated"})
}

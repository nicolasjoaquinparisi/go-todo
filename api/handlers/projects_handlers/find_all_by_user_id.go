package projects_handlers

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/models"
	"go-todo/api/repositories/projects_repository"
	"net/http"
)

func FindAllByUserId(c *gin.Context) {
	user, _ := c.Get("user")
	userId := user.(models.User).ID

	projects, err := projects_repository.FindAllByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Internal server error"})
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

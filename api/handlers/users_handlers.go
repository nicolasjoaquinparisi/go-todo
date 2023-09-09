package handlers

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/models"
	responses "go-todo/api/responses/users_response"
	"net/http"
)

func Me(c *gin.Context) {
	user, _ := c.Get("user")

	userResponse := responses.MeHandlerResponse{
		Email: user.(models.User).Email,
	}

	c.JSON(http.StatusOK, gin.H{"user": userResponse})
}

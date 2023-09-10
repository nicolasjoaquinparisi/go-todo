package users_handlers

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/models"
	responses "go-todo/api/responses/users_handlers_responses"
	"net/http"
)

func Me(c *gin.Context) {
	user, _ := c.Get("user")

	userResponse := responses.MeHandlerResponse{
		Email:     user.(models.User).Email,
		FirstName: user.(models.User).FirstName,
		LastName:  user.(models.User).LastName,
	}

	c.JSON(http.StatusOK, gin.H{"user": userResponse})
}

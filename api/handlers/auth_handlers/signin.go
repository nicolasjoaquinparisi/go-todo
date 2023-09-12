package auth_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go-todo/api/repositories/users_repository"
	"go-todo/api/utils/requests/auth_requests"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func SignIn(c *gin.Context) {
	var body auth_requests.SignInBodyStruct

	// map request body into body struct
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "Invalid payload"})
		return
	}

	// validate body
	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find user
	user, _ := users_repository.FindByEmail(body.Email)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"description": "Invalid credentials"})
		return
	}

	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"description": "Invalid credentials"})
		return
	}

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// sign and get the complete encoded token as a string using the secret
	var tokenString string
	tokenString, err = token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Internal server error"})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

package auth_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go-todo/api/repositories/users_repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func SignIn(c *gin.Context) {
	// get request body
	var body struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}

	var validate *validator.Validate
	var err error

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
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Internal server error"})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

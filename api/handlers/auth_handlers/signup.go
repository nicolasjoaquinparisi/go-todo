package auth_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-todo/api/repositories/users_repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required,min=6"`
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
	}

	var validate *validator.Validate
	var err error
	var hash []byte

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
	existUser, _ := users_repository.FindByEmail(body.Email)
	if existUser != nil {
		c.JSON(http.StatusConflict, gin.H{"description": "Email already in use"})
		return
	}

	// hash password
	hash, err = bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Internal server error"})
	}

	// create user
	_, err = users_repository.Create(body.Email, string(hash), body.FirstName, body.LastName)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"description": "Failed to create user"})
		return
	}

	// return response
	c.JSON(http.StatusCreated, gin.H{"description": "User created"})
}

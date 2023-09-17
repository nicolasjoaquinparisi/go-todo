package auth_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-todo/api/repositories/users_repository"
	"go-todo/api/utils/requests/auth_requests"
	"go-todo/api/utils/responses/auth_responses"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignUp(c *gin.Context) {
	fmt.Println("Invoked - Auth Handler SignUp")

	var body auth_requests.SignUpBodyStruct

	// map request body into body struct
	fmt.Println("Auth Handler SignUp - Validating request body struct")

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "Invalid payload"})
		return
	}

	// validate body
	fmt.Println("Auth Handler SignUp - Validating Request Body")

	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate duplicates
	fmt.Println("Auth Handler SignUp - Validating duplicates")

	existUser, _ := users_repository.FindByEmail(body.Email)
	if existUser != nil {
		c.JSON(http.StatusConflict, gin.H{"description": "Email already in use"})
		return
	}

	// hash password
	fmt.Println("Auth Handler SignUp - Hashing password")

	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Internal server error"})
	}

	// create user
	fmt.Println("Auth Handler SignUp - Creating user")

	user, err := users_repository.Create(body.Email, string(hash), body.FirstName, body.LastName)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"description": "Failed to create user"})
		return
	}

	// format response
	userResponse := auth_responses.SignupHandlerResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// return response
	c.JSON(http.StatusCreated, gin.H{"user": userResponse})
}

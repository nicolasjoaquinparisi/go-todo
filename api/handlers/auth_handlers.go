package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go-todo/api/models"
	"go-todo/database"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

var validate *validator.Validate
var err error
var hash []byte

func SignUp(c *gin.Context) {
	var body struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}

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
	var existUser models.User
	database.DB.First(&existUser, "email = ?", body.Email)
	if existUser.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"description": "Email already in use"})
		return
	}

	// hash password
	hash, err = bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Internal server error"})
		return
	}

	// create user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"description": "Failed to create user"})
		return
	}

	// return response
	c.JSON(http.StatusCreated, gin.H{"description": "User created"})
}

func SignIn(c *gin.Context) {
	// get request body
	var body struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}

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
	var user models.User
	database.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
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
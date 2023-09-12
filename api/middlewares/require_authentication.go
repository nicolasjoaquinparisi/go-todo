package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-todo/api/models"
	"go-todo/database"
	"net/http"
	"os"
	"strings"
	"time"
)

func RequireAuthentication(c *gin.Context) {
	headerAuthorization := c.GetHeader("Authorization")

	// validate if Authorization header is present in the request
	if headerAuthorization == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"description": "Missing Authorization header"})
		c.Abort()
		return
	}

	// get tokenString
	tokenString := strings.Split(headerAuthorization, " ")[1]
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		c.Abort()
		return
	}

	// decode jwt
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"description": "Token expired"})
		}

		// find user with token sub
		var user models.User
		database.Instance.First(&user, claims["sub"])

		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"description": "User not found"})
			c.Abort()
			return
		}

		// attach user to request
		c.Set("user", user)

		// continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}
}

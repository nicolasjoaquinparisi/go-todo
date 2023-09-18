package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go-todo/api/repositories/users_repository"
	"net/http"
	"os"
	"strings"
	"time"
)

func RequireAuthentication(c *gin.Context) {
	fmt.Println("Invoked: Middleware Authentication")

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
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"description": "Token expired"})
			c.Abort()
			return
		}

		// find user with token sub
		subClaim, _ := claims["sub"].(string)
		userId, _ := uuid.Parse(subClaim)
		user, _ := users_repository.FindById(userId)

		if user == nil {
			c.JSON(http.StatusForbidden, gin.H{"description": "Forbidden"})
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

package _test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go-todo/_test"
	"go-todo/api/handlers/auth_handlers"
	"go-todo/api/utils/requests/auth_requests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUp(t *testing.T) {
	_test.Setup()

	router := gin.New()
	router.POST("/api/auth/signup", auth_handlers.SignUp)

	// Create http test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Given
	requestBody := auth_requests.SignUpBodyStruct{
		Email:     "test@example.com",
		Password:  "password123",
		FirstName: "John",
		LastName:  "Doe",
	}

	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to convert struct into JSON: %v", err)
	}

	c.Request, _ = http.NewRequest("POST", "/api/auth/signup", bytes.NewReader(requestBodyJSON))
	c.Request.Header.Set("Content-Type", "application/json")
	auth_handlers.SignUp(c)

	// Then
	assert.Equal(t, 201, w.Code)
}

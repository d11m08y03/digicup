package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/d11m08y03/digicup/config"
	"github.com/d11m08y03/digicup/controllers"
	"github.com/d11m08y03/digicup/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	config.InitTestDB()
	router := gin.Default()
	router.POST("/register", controllers.Register)

	jsonBody := []byte(`{"name":"Alice","age":30,"email":"alice@example.com","password":"securepassword"}`)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "User registered successfully!")

	users, err := models.GetAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "Alice", users[0].Name)
}

func TestGetUsers(t *testing.T) {
	config.InitTestDB()
	router := gin.Default()
	router.GET("/users", controllers.GetUsers)

	err := models.CreateUser(models.User{ID: 1, Name: "John", Age: 25, Email: "john@example.com", Password: "hashedpassword"})
	assert.NoError(t, err)

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John")
}

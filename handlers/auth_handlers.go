package handlers

import (
	"go-authentication-api/models"
	"go-authentication-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// In-memory user store (username -> User)
var userStore = map[string]models.User{
	"user": {ID: 1, Username: "user", Password: "password"},
}

func Login(c *gin.Context) {
	var inputUser models.User

	// Check user credentials and generate a JWT token
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Check if credentials are valid
	storedUser, exists := userStore[inputUser.Username]
	if !exists || storedUser.Password != inputUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a JWT token
	token, err := utils.GenerateToken(storedUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Function for registering a new user
func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Check if user already exists
	if _, exists := userStore[user.Username]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// Assign ID (simple auto-increment simulation)
	user.ID = uint(len(userStore) + 1)

	// Save to store
	userStore[user.Username] = user

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

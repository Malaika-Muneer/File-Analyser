package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/middleware"
	"github.com/malaika-muneer/File-Analyser/models"
	"github.com/malaika-muneer/File-Analyser/service"
)

var jwtSecret = []byte("your-secret-key")

// SignInHandler handles the sign-in process
func SignInHandler(c *gin.Context) {
	var signInData models.SignIn

	if err := c.ShouldBindJSON(&signInData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Call service layer for authentication
	user, err := service.AuthenticateUser(signInData.Username, signInData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := middleware.GenerateJWT(user.Username, user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}

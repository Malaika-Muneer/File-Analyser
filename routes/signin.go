package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/middleware"
	"github.com/malaika-muneer/File-Analyser/models"
)

func (r *Router) SignInHandler(c *gin.Context) {
	var signInData models.SignIn
	if err := c.ShouldBindJSON(&signInData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	user, err := r.userService.AuthenticateUser(signInData.Username, signInData.Password) // Call service layer for authentication
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	token, err := middleware.GenerateJWT(user.Username, user.Id) // Generate JWT token
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}

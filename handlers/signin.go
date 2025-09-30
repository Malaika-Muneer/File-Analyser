package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/db"
	"github.com/malaika-muneer/File-Analyser/middleware"
	"github.com/malaika-muneer/File-Analyser/models"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your-secret-key")

// SignInHandler handles the sign-in process
func SignInHandler(c *gin.Context) {
	var signInData models.SignIn

	if err := c.ShouldBindJSON(&signInData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var storedUser models.User
	query := "SELECT id, username, password FROM users WHERE username = ? OR email = ?"
	err := db.DB.QueryRow(query, signInData.Username, signInData.Username).
		Scan(&storedUser.Id, &storedUser.Username, &storedUser.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(signInData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// ✅ generate token with string ID
	token, err := middleware.GenerateJWT(storedUser.Username, storedUser.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   token,
	})
}

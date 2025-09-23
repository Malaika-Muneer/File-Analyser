package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/models"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your-secret-key") // Change this to your secret key for signing tokens

// SignInHandler handles the sign-in process
func SignInHandler(c *gin.Context) {
	var signInData models.SignIn

	if err := c.ShouldBindJSON(&signInData); err != nil {
		log.Printf("Error parsing request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Body request"})
		return
	}

	// Query the database for the user by username/email
	var storedUser models.User
	query := "SELECT username, password FROM users WHERE username = ? OR email = ?"
	err := DbConnection.DB.QueryRow(query, signInData.Username, signInData.Username).Scan(&storedUser.Username, &storedUser.Password)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Printf("Invalid credentials for username/email: %s", signInData.Username)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		} else {
			log.Printf("Database error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Querynig databases"})
		}
		return
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(signInData.Password)); err != nil {
		log.Printf("Password mismatch for user: %s", signInData.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}

	// Log successful login
	log.Printf("User %s signed in successfully", signInData.Username)

	// Generate JWT token
	token, err := generateJWT(storedUser.Username)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating Token"})
		return
	}

	// Respond with token
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   token,
	})

}

// generateJWT creates a JWT token for the user
func generateJWT(username string) (string, error) {
	// Set expiration time for the token (e.g., 24 hours)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the claims for the token
	claims := &jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
	}

	// Create the JWT token with the claims and the signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

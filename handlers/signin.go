package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/malaika-muneer/File-Analyser/DbConnection"
	"github.com/malaika-muneer/File-Analyser/models"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your-secret-key") // Change this to your secret key for signing tokens

// SignInHandler handles the sign-in process
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Invalid request method: %s", r.Method)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON request body into the SignIn struct
	var signInData models.SignIn
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&signInData)
	if err != nil {
		log.Printf("Error parsing request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Query the database for the user by username/email
	var storedUser models.User
	query := "SELECT username, password FROM users WHERE username = ? OR email = ?"
	err = DbConnection.DB.QueryRow(query, signInData.Username, signInData.Username).Scan(&storedUser.Username, &storedUser.Password)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Printf("Invalid credentials for username/email: %s", signInData.Username)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			log.Printf("Database error: %v", err)
			http.Error(w, "Error querying database", http.StatusInternalServerError)
		}
		return
	}

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(signInData.Password))
	if err != nil {
		log.Printf("Password mismatch for user: %s", signInData.Username)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Log successful login
	log.Printf("User %s signed in successfully", signInData.Username)

	// Generate JWT token
	token, err := generateJWT(storedUser.Username)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Respond with the token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
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

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

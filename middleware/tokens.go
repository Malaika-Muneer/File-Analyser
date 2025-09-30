package middleware

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ✅ Fixed generateJWT
func GenerateJWT(username string, id int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"username": username,
		"id":       fmt.Sprintf("%d", id), // store ID as string
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

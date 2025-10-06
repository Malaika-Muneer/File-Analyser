package service

import (
	"database/sql"
	"errors"

	"github.com/malaika-muneer/File-Analyser/db"
	"github.com/malaika-muneer/File-Analyser/models"
	"golang.org/x/crypto/bcrypt"
)

// SignupUser handles the signup logic
func SignupUser(user models.User) error {
	// Basic input validation
	if user.Username == "" || user.Password == "" {
		return errors.New("username or password cannot be empty")
	}

	// Hash password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing password")
	}
	user.Password = string(hashedPassword)

	if err := db.InsertUser(user); err != nil {
		return err
	}

	return nil
}

// AuthenticateUser handles user signin (used by SignInHandler)
func AuthenticateUser(usernameOrEmail, password string) (*models.User, error) {
	database := db.GetDB()

	var storedUser models.User
	query := "SELECT id, username, password FROM users WHERE username = ? OR email = ?"
	err := database.QueryRow(query, usernameOrEmail, usernameOrEmail).
		Scan(&storedUser.Id, &storedUser.Username, &storedUser.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Compare hashed password with the entered password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return &storedUser, nil
}

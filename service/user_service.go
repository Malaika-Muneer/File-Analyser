package service

import (
	"database/sql"
	"errors"

	"github.com/malaika-muneer/File-Analyser/db"
	"github.com/malaika-muneer/File-Analyser/models"
	"golang.org/x/crypto/bcrypt"
)

// SignupUser handles the signup logic
func (s *UserServiceImpl) SignupUser(user models.User) error {

	if user.Username == "" || user.Password == "" {
		return errors.New("username or password cannot be empty")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing password")
	}
	user.Password = string(hashedPassword)

	if err := s.Dao.InsertUser(user); err != nil {
		return err
	}

	return nil
}
func (s *UserServiceImpl) AuthenticateUser(usernameOrEmail, password string) (*models.User, error) {
	database := db.ConnectDB()

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

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return &storedUser, nil
}

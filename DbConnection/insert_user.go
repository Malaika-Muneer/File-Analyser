package dbConnection

import (
	"log"

	"github.com/malaika-muneer/File-Analyser/models"
)

// InsertUser inserts a new user into the database
func InsertUser(user models.User) error {
	query := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`
	_, err := DB.Exec(query, user.Username, user.Password, user.Email)
	if err != nil {
		log.Println("Error inserting user:", err)
		return err
	}
	return nil
}

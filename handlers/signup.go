package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	dbConnection "github.com/malaika-muneer/File-Analyser/DbConnection"
	"github.com/malaika-muneer/File-Analyser/models"
	"golang.org/x/crypto/bcrypt"
)

// signupHandler handles the sign-up process
func SignupHandler(c *gin.Context) {

	var user models.User
	//binding JSON request to struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Inavlid Request Body "})
		return
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error hashing password"})
		return
	}

	user.Password = string(hashedPassword)

	// Insert the user data into the database

	if err := dbConnection.InsertUser(user); err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error inserting user into database"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message ": "user created succesfully"})
	log.Printf("user created succesfully and stored in database")
}

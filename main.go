package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/db"
	"github.com/malaika-muneer/File-Analyser/routes"
)

func main() {
	// Connect to the database
	db.ConnectDB()
	if db.GetDB() == nil {
		log.Fatal("Database connection failed — check your .env configuration")
	}
	fmt.Println("Database connection established successfully")

	// Create uploads folder if it doesn’t exist
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		err := os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create uploads directory: %v", err)
		}
		fmt.Println("Created uploads directory")
	}

	// Initialize Gin router
	r := gin.Default()

	// Setup routes (signup, login, file upload, etc.)
	routes.SetupRoutes(r)

	// Start server
	fmt.Println("Server started at http://localhost:8005")
	if err := r.Run(":8005"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

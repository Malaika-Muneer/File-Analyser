package main

import (
	"fmt"
	"log"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/dbConnection"
	routes "github.com/malaika-muneer/File-Analyser/routes"
)

func main() {

	dbConnection.ConnectDB()

	// Create uploads folder if not exists
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		err := os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create uploads directory: %v", err)
		}
	}

	// Initialize Gin router
	r := gin.Default()
	// Public routes
	routes.SetupRoutes(r)
	fmt.Println("Server started at http://localhost:8005")
	r.Run(":8005")

}

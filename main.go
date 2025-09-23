package main

import (
	"fmt"
	"log"

	"os"

	"github.com/gin-gonic/gin"
	dbConnection "github.com/malaika-muneer/File-Analyser/DbConnection"
	"github.com/malaika-muneer/File-Analyser/Routes"
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
	Routes.SetupRoutes(r)
	fmt.Println("Server started at http://localhost:8005")
	r.Run(":8005")

}

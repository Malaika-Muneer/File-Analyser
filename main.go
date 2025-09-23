package main

import (
	"fmt"
	"log"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/DbConnection"
	"github.com/malaika-muneer/File-Analyser/handlers"
	"github.com/malaika-muneer/File-Analyser/middleware"
)

func main() {

	DbConnection.ConnectDB()

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
	r.POST("/upload", handlers.UploadFile)
	r.POST("/signup", handlers.SignupHandler)
	r.POST("/signin", handlers.SignInHandler)

	r.GET("/protected", middleware.TokenValidationMiddleware(), func(c *gin.Context) {
		c.String(200, "This is a protected route, you are authorized!")
	})

	fmt.Println("Server started at http://localhost:8005")
	r.Run(":8005")

}

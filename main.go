package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/malaika-muneer/File-Analyser/db"
	_ "github.com/malaika-muneer/File-Analyser/docs"
	"github.com/malaika-muneer/File-Analyser/routes"
	"github.com/malaika-muneer/File-Analyser/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title File Analyser
// @version 2.0
// @description This is a file analyser application API built with Go and Gin.
// @host localhost:8005
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	conn := db.ConnectDB()
	dao := db.NewDao(conn)
	fmt.Println("Database connection established successfully")

	// Create uploads folder if it doesn’t exist
	// if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
	// 	err := os.Mkdir("./uploads", os.ModePerm)
	// if err != nil {
	// 		log.Fatalf("Failed to create uploads directory: %v", err)
	// 	}
	// 	fmt.Println("Created uploads directory")
	// }

	// Initialize Gin router
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	userService := service.NewUserService(dao)
	router := routes.NewRouter(userService)
	// Setup routes (signup, login, file upload, etc.)
	router.SetupRoutes(r)
	// Start server
	fmt.Println("Server started at http://localhost:8005")
	if err := r.Run(":8005"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

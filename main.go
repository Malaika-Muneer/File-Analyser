package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/db"
	"github.com/malaika-muneer/File-Analyser/db/mongodb" // MongoDB integration
	_ "github.com/malaika-muneer/File-Analyser/docs"     // Swagger docs
	"github.com/malaika-muneer/File-Analyser/routes"
	"github.com/malaika-muneer/File-Analyser/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//
// @title           File Analyser API
// @version         1.0
// @description     This project analyses text files by splitting them into chunks, counting characters, vowels, consonants, etc. It demonstrates integration of MySQL for user management and MongoDB for analysis storage.
//
// @contact.name    Malaika Muneer
// @contact.url     https://github.com/malaika-muneer
// @contact.email   malaikamuneer@gmail.com
//
// @license.name    MIT License
// @license.url     https://opensource.org/licenses/MIT
//
// @BasePath  /
//
// @host       localhost:8005
// @schemes    http
//

func main() {
	// ---- Connect MySQL (Users) ----
	// MySQL handles user authentication and registration.
	conn := db.ConnectDB()
	dao := db.NewDao(conn)
	fmt.Println("MySQL connection established successfully")

	// ---- Connect MongoDB (Analysis Results) ----
	// MongoDB stores file analysis data for each uploaded chunk.
	mongoColl := mongodb.ConnectMongo()
	mongoDAO := mongodb.NewMongo(mongoColl)
	fmt.Println("MongoDB connection established successfully")

	// ---- Initialize Services ----
	// Services are the middle layer between routes and database logic.
	userService := service.NewUserService(dao)          // handles MySQL (users)
	uploadService := service.NewUploadService(mongoDAO) // handles MongoDB (analysis)

	// ---- Setup Gin Router ----
	r := gin.Default()

	// Swagger UI setup — visit http://localhost:8005/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ---- Register Routes ----
	router := routes.NewRouter(userService, uploadService)
	router.SetupRoutes(r)

	// ---- Serve Frontend Static Files ----
	// These routes serve your HTML frontend pages.
	r.Static("/frontend", "./frontend")
	r.GET("/", func(c *gin.Context) { c.File("./frontend/index.html") })
	r.GET("/signup.html", func(c *gin.Context) { c.File("./frontend/signup.html") })
	r.GET("/signin.html", func(c *gin.Context) { c.File("./frontend/signin.html") })
	r.GET("/upload.html", func(c *gin.Context) { c.File("./frontend/upload.html") })

	// ---- Start Server ----
	fmt.Println("🚀 Server running at: http://localhost:8005")
	if err := r.Run(":8005"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

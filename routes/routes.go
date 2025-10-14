package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/malaika-muneer/File-Analyser/middleware"
)

func (r *Router) SetupRoutes(gin *gin.Engine) {

	gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5500"}, // your frontend address
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	gin.POST("/signup", r.SignupHandler)
	gin.POST("/signin", r.SignInHandler)
	auth := gin.Group("/")
	auth.Use(middleware.TokenValidationMiddleware())
	auth.POST("/upload", r.UploadFilehandler)

}

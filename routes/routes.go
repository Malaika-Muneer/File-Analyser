package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/handlers"
	"github.com/malaika-muneer/File-Analyser/middleware"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/signup", handlers.SignupHandler)
	r.POST("/signin", handlers.SignInHandler)

	auth := r.Group("/")
	auth.Use(middleware.TokenValidationMiddleware())

	auth.POST("/upload", handlers.UploadFile)

}

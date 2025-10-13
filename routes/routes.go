package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/malaika-muneer/File-Analyser/middleware"
)

func (r *Router) SetupRoutes(gin *gin.Engine) {
	gin.POST("/signup", r.SignupHandler)
	gin.POST("/signin", r.SignInHandler)
	auth := gin.Group("/")
	auth.Use(middleware.TokenValidationMiddleware())
	auth.POST("/upload", r.UploadFilehandler)

}

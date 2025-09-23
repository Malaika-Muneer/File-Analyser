package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/handlers"
	"github.com/malaika-muneer/File-Analyser/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/upload", handlers.UploadFile)
	r.POST("/signup", handlers.SignupHandler)
	r.POST("/signin", handlers.SignInHandler)

	r.GET("/protected", middleware.TokenValidationMiddleware(), func(c *gin.Context) {
		c.String(200, "This is a protected route, you are authorized!")
	})

}

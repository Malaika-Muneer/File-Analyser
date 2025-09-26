package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/handlers"
	"github.com/malaika-muneer/File-Analyser/middleware"
)

func SetupRoutes(r *gin.Engine) {
	//r.POST("/upload", handlers.UploadFile)
	//these are public routes
	r.POST("/signup", handlers.SignupHandler)
	r.POST("/signin", handlers.SignInHandler)

	auth := r.Group("/")
	auth.Use(middleware.TokenValidationMiddleware())

	auth.POST("/upload", handlers.UploadFile)

	auth.GET("/protected", func(c *gin.Context) {
		c.String(200, "This is a protected route, you are authorized!")
	})

}

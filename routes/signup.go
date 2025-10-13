package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/models"
)

func (r *Router) SignupHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := r.userService.SignupUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/models"
)

// @Summary      User SignUp
// @Description  Register a new user
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        user  body  models.User  true  "User Data"
// @Success      201  {object}  map[string]interface{}  "User created successfully"
// @Failure      400  {object}  map[string]interface{}  "Invalid request"
// @Failure      500  {object}  map[string]interface{}  "Internal Server Error"
// @Router       /signup [post]
func (r *Router) SignupHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := r.UserService.SignupUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

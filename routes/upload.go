package routes

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadFileHandler godoc
// @Summary      Upload and analyze a file
// @Description  Allows an authenticated user to upload a file for analysis. The file is split into chunks, and each chunk is analyzed separately.
// @Tags         File
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "File to be uploaded"
// @Success      200  {object}  map[string]interface{}  "File uploaded successfully with chunked analysis"
// @Failure      400  {object}  map[string]string       "Failed to read uploaded file"
// @Failure      401  {object}  map[string]string       "User not found or unauthorized"
// @Failure      500  {object}  map[string]string       "Internal server error"
// @Security     BearerAuth
// @Router       /upload [post]
func (r *Router) UploadFilehandler(c *gin.Context) {
	log.Println("Upload endpoint hit")

	username, userExists := c.Get("username")
	id, idExists := c.Get("id")
	if !userExists || !idExists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read uploaded file"})
		return
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
		return
	}

	analyses, err := r.userService.UploadFile(fileContent, username.(string), id.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error analyzing file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"chunks":  analyses, // ✅ Key changed to "chunks"
	})
}

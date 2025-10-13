package routes

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) UploadFilehandler(c *gin.Context) {
	log.Println("Upload endpoint hit")
	// Get username from middleware (set in context)
	// username, exists := c.Get("username")
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	// Get the uploaded_file
	file, _, err := c.Request.FormFile("uploadedFile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read Uploaded file"})
		return
	}
	defer file.Close()
	// Read the entire content of the file
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file content."})
		return
	}
	data, err := r.userService.UploadFile(fileContent, "username", id.(int))
	log.Println("File content size:", len(fileContent))
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"data":    data,
	})
}

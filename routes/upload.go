package routes

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/malaika-muneer/File-Analyser/models"
)

// UploadFileHandler godoc
// @Summary      Upload and analyze a file
// @Description  Allows an authenticated user to upload a file for analysis. The file is split into chunks, and each chunk is analyzed separately.
// @Tags         File
// @Accept       multipart/form-data
// @Produce      json
// @Param        file       formData  file  true  "File to be uploaded"
// @Param        numChunks  formData  int   true  "Number of chunks to divide file into"
// @Success      200  {object}  map[string]interface{}  "File uploaded successfully with chunked analysis"
// @Failure      400  {object}  map[string]string       "Failed to read uploaded file or invalid chunk number"
// @Failure      401  {object}  map[string]string       "User not found or unauthorized"
// @Failure      500  {object}  map[string]string       "Internal server error"
// @Security     BearerAuth
// @Router       /upload [post]
func (r *Router) UploadFilehandler(c *gin.Context) {
	log.Println("Upload endpoint hit")

	// Get user info from context
	username, userExists := c.Get("username")
	id, idExists := c.Get("id")
	if !userExists || !idExists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	// Read uploaded file
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

	// Read number of chunks from form
	numChunksStr := c.PostForm("numChunks")
	numChunks, err := strconv.Atoi(numChunksStr)
	if err != nil || numChunks < 1 {
		numChunks = 1 // fallback to 1 chunk if invalid
	}

	// Call service with numChunks
	resultMap, err := r.userService.UploadFile(fileContent, username.(string), id.(int), numChunks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error analyzing file"})
		return
	}

	// Extract sequential and concurrent analyses and execution times
	sequential, _ := resultMap["sequential"].([]models.FileAnalysis)
	concurrent, _ := resultMap["concurrent"].([]models.FileAnalysis)
	timeSeq, _ := resultMap["timeSeq"].(int64)
	timeCon, _ := resultMap["timeCon"].(int64)

	c.JSON(http.StatusOK, gin.H{
		"message":        "File uploaded successfully",
		"sequential":     sequential,
		"concurrent":     concurrent,
		"timeSequential": timeSeq, // in milliseconds
		"timeConcurrent": timeCon, // in milliseconds
	})
}

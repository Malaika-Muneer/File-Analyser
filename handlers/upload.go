package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	dbConnection "github.com/malaika-muneer/File-Analyser/DbConnection"
	"github.com/malaika-muneer/File-Analyser/models"
)

// Handle file upload, analyze it, and return JSON response
func UploadFile(c *gin.Context) {
	log.Println("Upload endpoint hit")

	// Get the uploaded file
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

	// Create a channel for the analysis result
	analysisCh := make(chan models.FileAnalysis)

	// Use a goroutine to perform the file analysis concurrently
	go analyzeFileConcurrently(fileContent, analysisCh)

	// Wait for the result from the channel
	analysis := <-analysisCh

	dbConnection.InsertAnalysisData(analysis)

	c.JSON(http.StatusOK, analysis)
}

// Function to analyze the file content concurrently
func analyzeFileConcurrently(content []byte, ch chan models.FileAnalysis) {
	// Perform analysis on the file content
	analysis := analyzeFile(content)

	// Send the result to the channel
	ch <- analysis
}

// Function to analyze the file content and return the analysis
func analyzeFile(content []byte) models.FileAnalysis {
	var analysis models.FileAnalysis

	for _, char := range content {
		runeChar := rune(char)

		// Check for spaces
		if unicode.IsSpace(runeChar) {
			analysis.Spaces++
		}

		// Check if the character is a letter
		if unicode.IsLetter(runeChar) {
			analysis.Letters++
			// Check for uppercase and lowercase letters
			if unicode.IsUpper(runeChar) {
				analysis.UpperCase++
			} else if unicode.IsLower(runeChar) {
				analysis.LowerCase++
			}
		}

		// Check if the character is a digit
		if unicode.IsDigit(runeChar) {
			analysis.Digits++
		}

		// Check if the character is a vowel (both uppercase and lowercase)
		if isVowel(runeChar) {
			analysis.Vowels++
		} else if unicode.IsLetter(runeChar) {
			analysis.Consonants++
		}

		// Check for special characters (non-alphanumeric)
		if !unicode.IsLetter(runeChar) && !unicode.IsDigit(runeChar) && !unicode.IsSpace(runeChar) {
			analysis.SpecialChars++
		}
	}

	// Total character count (including spaces)
	analysis.TotalChars = len(content)

	return analysis
}

// Helper function to check if a character is a vowel
func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, r)
}

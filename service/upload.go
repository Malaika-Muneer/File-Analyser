package service

import (
	"strings"
	"unicode"

	"github.com/malaika-muneer/File-Analyser/models"
)

// UploadFilehandler godoc
// @Summary      Upload and analyze a file
// @Description  Uploads a file for analysis (requires authentication)
// @Tags         File
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "File to upload"
// @Security     BearerAuth
// @Success      200  {object}  models.FileAnalysis
// @Failure      400  {string}  string  "Bad request"
// @Router       /upload [post]
// Handle file upload, analyze it, and return JSON response
func (s *UserServiceImpl) UploadFile(fileContent []byte, Username string, id int) (models.FileAnalysis, error) {
	analysisCh := make(chan models.FileAnalysis)
	go analyzeFileConcurrently(fileContent, analysisCh)
	analysis := <-analysisCh
	analysis.Username = Username
	analysis.Id = id

	err := s.Dao.InsertAnalysisData(analysis)
	if err != nil {
		return models.FileAnalysis{}, err
	}

	return analysis, nil
}

// Function to analyze the file content concurrently
func analyzeFileConcurrently(content []byte, ch chan models.FileAnalysis) {
	analysis := analyzeFile(content)
	ch <- analysis
}

// Function to analyze the file content and return the analysis
func analyzeFile(content []byte) models.FileAnalysis {
	var analysis models.FileAnalysis

	for _, char := range content {
		runeChar := rune(char)

		if unicode.IsSpace(runeChar) {
			analysis.Spaces++
		}

		if unicode.IsLetter(runeChar) {
			analysis.Letters++
			if unicode.IsUpper(runeChar) {
				analysis.UpperCase++
			} else if unicode.IsLower(runeChar) {
				analysis.LowerCase++
			}
		}
		if unicode.IsDigit(runeChar) {
			analysis.Digits++
		}
		if isVowel(runeChar) {
			analysis.Vowels++
		} else if unicode.IsLetter(runeChar) {
			analysis.Consonants++
		}
		if !unicode.IsLetter(runeChar) && !unicode.IsDigit(runeChar) && !unicode.IsSpace(runeChar) {
			analysis.SpecialChars++
		}
	}

	analysis.TotalChars = len(content)

	return analysis
}

func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, r)
}

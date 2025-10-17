package service

import (
	"log"
	"strings"
	"unicode"

	"github.com/malaika-muneer/File-Analyser/models"
)

func (s *UserServiceImpl) UploadFile(fileContent []byte, username string, id int) ([]models.FileAnalysis, error) {
	chunkSize := 1024 // 1KB per chunk
	var analyses []models.FileAnalysis

	totalChunks := (len(fileContent) + chunkSize - 1) / chunkSize

	for i := 0; i < totalChunks; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(fileContent) {
			end = len(fileContent)
		}

		chunk := fileContent[start:end]
		analysis := analyzeFile(chunk)

		analysis.Username = username
		analysis.Id = id
		analysis.ChunkNumber = i + 1
		log.Println("Inserting for user ID:", id)

		if err := s.Dao.InsertAnalysisData(analysis); err != nil {
			return nil, err
		}

		analyses = append(analyses, analysis)
	}

	return analyses, nil
}

func analyzeFile(content []byte) models.FileAnalysis {
	var analysis models.FileAnalysis

	for _, char := range content {
		r := rune(char)

		if unicode.IsSpace(r) {
			analysis.Spaces++
		}

		if unicode.IsLetter(r) {
			analysis.Letters++
			if unicode.IsUpper(r) {
				analysis.UpperCase++
			} else {
				analysis.LowerCase++
			}
		}

		if unicode.IsDigit(r) {
			analysis.Digits++
		}

		if isVowel(r) {
			analysis.Vowels++
		} else if unicode.IsLetter(r) {
			analysis.Consonants++
		}

		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
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

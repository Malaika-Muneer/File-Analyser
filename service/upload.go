package service

import (
	"log"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/malaika-muneer/File-Analyser/models"
)

// UploadFile analyzes the file in both sequential and concurrent ways
func (s *UserServiceImpl) UploadFile(fileContent []byte, username string, id int, numChunks int) (map[string]interface{}, error) {
	if numChunks < 1 {
		numChunks = 1
	}

	chunkSize := (len(fileContent) + numChunks - 1) / numChunks
	var sequentialAnalyses []models.FileAnalysis

	// ----------- Case 1: Sequential -----------
	startSeq := time.Now()
	for i := 0; i < numChunks; i++ {
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

		if err := s.Dao.InsertAnalysisData(analysis); err != nil {
			return nil, err
		}

		sequentialAnalyses = append(sequentialAnalyses, analysis)
	}
	execTimeSequential := time.Since(startSeq).Milliseconds() // in milliseconds

	// ----------- Case 2: Concurrent using Goroutines -----------
	startConcurrent := time.Now()
	concurrentAnalyses := make([]models.FileAnalysis, numChunks)
	var wg sync.WaitGroup
	wg.Add(numChunks)

	for i := 0; i < numChunks; i++ {
		go func(i int) {
			defer wg.Done()
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
			concurrentAnalyses[i] = analysis

			// Optional: Insert into DB sequentially
			if err := s.Dao.InsertAnalysisData(analysis); err != nil {
				log.Println("Error inserting analysis in goroutine:", err)
			}
		}(i)
	}

	wg.Wait()
	execTimeConcurrent := time.Since(startConcurrent).Milliseconds() // in milliseconds

	// Return both analyses and execution times
	return map[string]interface{}{
		"sequential": sequentialAnalyses,
		"concurrent": concurrentAnalyses,
		"timeSeq":    execTimeSequential,
		"timeCon":    execTimeConcurrent,
	}, nil
}

// analyzeFile remains same
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

	analysis.TotalChars = analysis.Letters + analysis.Digits + analysis.Spaces + analysis.SpecialChars
	return analysis
}

// isVowel remains same
func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, r)
}

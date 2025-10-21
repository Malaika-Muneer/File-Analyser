package service

import (
	"log"
	"runtime"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/malaika-muneer/File-Analyser/models"
)

func (s *UploadService) UploadFile(fileContent []byte, username string, id int, numChunks int) (map[string]interface{}, error) {
	if numChunks < 1 {
		numChunks = 1
	}

	chunkSize := (len(fileContent) + numChunks - 1) / numChunks
	var sequentialAnalyses []models.FileAnalysis

	// ----------- Sequential Analysis -----------
	startSeq := time.Now()
	for i := 0; i < numChunks; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(fileContent) {
			end = len(fileContent)
		}

		var a models.FileAnalysis
		for _, char := range fileContent[start:end] {
			r := rune(char)

			if unicode.IsSpace(r) {
				a.Spaces++
			}
			if unicode.IsLetter(r) {
				a.Letters++
				if unicode.IsUpper(r) {
					a.UpperCase++
				} else {
					a.LowerCase++
				}
			}
			if unicode.IsDigit(r) {
				a.Digits++
			}
			if strings.ContainsRune("aeiouAEIOU", r) {
				a.Vowels++
			} else if unicode.IsLetter(r) {
				a.Consonants++
			}
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
				a.SpecialChars++
			}
		}

		a.TotalChars = a.Letters + a.Digits + a.Spaces + a.SpecialChars
		a.Username = username
		a.Id = id
		a.ChunkNumber = i + 1

		// ✅ Store in MongoDB
		if err := s.MongoDAO.InsertAnalysisData(a); err != nil {
			return nil, err
		}

		sequentialAnalyses = append(sequentialAnalyses, a)
	}
	execTimeSequential := time.Since(startSeq).Milliseconds()

	// ----------- Concurrent Analysis -----------
	runtime.GOMAXPROCS(runtime.NumCPU())

	startConcurrent := time.Now()
	var wg sync.WaitGroup
	wg.Add(numChunks)

	results := make([]models.FileAnalysis, numChunks)

	for i := 0; i < numChunks; i++ {
		i := i
		go func() {
			defer wg.Done()
			start := i * chunkSize
			end := start + chunkSize
			if end > len(fileContent) {
				end = len(fileContent)
			}

			var a models.FileAnalysis
			for _, char := range fileContent[start:end] {
				r := rune(char)

				if unicode.IsSpace(r) {
					a.Spaces++
				}
				if unicode.IsLetter(r) {
					a.Letters++
					if unicode.IsUpper(r) {
						a.UpperCase++
					} else {
						a.LowerCase++
					}
				}
				if unicode.IsDigit(r) {
					a.Digits++
				}
				if strings.ContainsRune("aeiouAEIOU", r) {
					a.Vowels++
				} else if unicode.IsLetter(r) {
					a.Consonants++
				}
				if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
					a.SpecialChars++
				}
			}

			a.TotalChars = a.Letters + a.Digits + a.Spaces + a.SpecialChars
			a.Username = username
			a.Id = id
			a.ChunkNumber = i + 1

			results[i] = a
		}()
	}

	wg.Wait()

	// ✅ Insert concurrently calculated data into MongoDB
	for _, a := range results {
		if err := s.MongoDAO.InsertAnalysisData(a); err != nil {
			log.Println("Mongo insert error:", err)
		}
	}

	execTimeConcurrent := time.Since(startConcurrent).Milliseconds()

	return map[string]interface{}{
		"sequential": sequentialAnalyses,
		"concurrent": results,
		"timeSeq":    execTimeSequential,
		"timeCon":    execTimeConcurrent,
	}, nil
}

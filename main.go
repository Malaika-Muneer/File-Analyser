package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
	"unicode"
)

// Main function to start the server and handle routes
func main() {
	http.HandleFunc("/", uploadForm)       // Form for uploading
	http.HandleFunc("/upload", uploadFile) // File upload handler
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Show the file upload form
func uploadForm(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
		<body>
			<h2>Upload a File for Analysis</h2>
			<form enctype="multipart/form-data" action="/upload" method="post">
				<input type="file" name="uploadedFile" />
				<input type="submit" value="Upload" />
			</form>
		</body>
		</html>
	`
	fmt.Fprint(w, html)
}

// Handle the file upload
func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read uploaded file
	file, header, err := r.FormFile("uploadedFile")
	if err != nil {
		http.Error(w, "Failed to read uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the file locally
	dst, err := os.Create("./uploads/" + header.Filename)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
		return
	}

	// Run analysis
	analysisResult := analyzeFile("./uploads/" + header.Filename)

	// Return analysis result to the user
	fmt.Fprintf(w, "<h3>File uploaded and analyzed successfully.</h3><pre>%s</pre>", analysisResult)
}

// Analyze the file and count characters
func analyzeFile(filePath string) string {
	start := time.Now() // Start timing

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Use a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Channel to receive the counts from each goroutine
	ch := make(chan map[string]int, 100) // Buffered channel to hold counts

	var vowels, consonants, digits, specialChars, letters int

	// Read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wg.Add(1)                               // Increment wait group for each goroutine
		go processLine(scanner.Text(), &wg, ch) // Start a goroutine for each line
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Collect the counts from the channel
	for counts := range ch {
		vowels += counts["vowels"]
		consonants += counts["consonants"]
		digits += counts["digits"]
		specialChars += counts["specialChars"]
		letters += counts["letters"]
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Format the analysis result
	duration := time.Since(start) // End timing
	analysisResult := fmt.Sprintf(`
		File Analysis:
		Total Letters     : %d
		Total Vowels      : %d
		Total Consonants  : %d
		Total Digits      : %d
		Special Characters: %d
		Execution Time    : %.4f seconds
	`, letters, vowels, consonants, digits, specialChars, duration.Seconds())

	return analysisResult
}

// Process a line and count characters
func processLine(text string, wg *sync.WaitGroup, ch chan<- map[string]int) {
	defer wg.Done()

	// Local counts for this line
	localCounts := map[string]int{
		"vowels":       0,
		"consonants":   0,
		"digits":       0,
		"specialChars": 0,
		"letters":      0,
	}

	// Process each character in the line
	for _, ch := range text {
		// Fast path: letters
		if unicode.IsLetter(ch) {
			localCounts["letters"]++
			switch ch | 0x20 { // Bitwise OR for fast lowercase (ASCII only)
			case 'a', 'e', 'i', 'o', 'u':
				localCounts["vowels"]++
			default:
				localCounts["consonants"]++
			}
			continue
		}

		// Digits (ASCII only)
		if ch >= '0' && ch <= '9' {
			localCounts["digits"]++
			continue
		}

		// Space
		if unicode.IsSpace(ch) {
			continue
		}

		// Everything else is special
		localCounts["specialChars"]++
	}

	// Send the local counts to the channel
	ch <- localCounts
}

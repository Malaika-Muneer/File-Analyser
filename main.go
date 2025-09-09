package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"unicode"
)

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

func main() {
	start := time.Now() // Start timing

	filePath := `C:\Users\Dell\Documents\practice.txt`

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

	fmt.Println("File Analysis:")
	fmt.Printf("Total Letters     : %d\n", letters)
	fmt.Printf("Total Vowels      : %d\n", vowels)
	fmt.Printf("Total Consonants  : %d\n", consonants)
	fmt.Printf("Total Digits      : %d\n", digits)
	fmt.Printf("Special Characters: %d\n", specialChars)

	duration := time.Since(start) // End timing
	fmt.Printf("Execution Time     : %.4f seconds\n", duration.Seconds())
}

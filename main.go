package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"unicode"
)

func main() {
	start := time.Now() //Start timing

	filePath := `C:\Users\Dell\Documents\practice.txt`

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var vowels, consonants, digits, specialChars, letters int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		for _, ch := range scanner.Text() {
			// Fast path: letters
			if unicode.IsLetter(ch) {
				letters++
				switch ch | 0x20 { // Bitwise OR for fast lowercase (ASCII only)
				case 'a', 'e', 'i', 'o', 'u':
					vowels++
				default:
					consonants++
				}
				continue
			}

			// Digits (ASCII only)
			if ch >= '0' && ch <= '9' {
				digits++
				continue
			}

			// Space
			if unicode.IsSpace(ch) {
				continue
			}

			// Everything else is special
			specialChars++
		}
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

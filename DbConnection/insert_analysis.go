package dbConnection

import (
	"log"

	"github.com/malaika-muneer/File-Analyser/models"
)

func InsertAnalysisData(analysis models.FileAnalysis) error {
	log.Printf("Database insert")
	query := `
    INSERT INTO analysis_results (vowels, consonants, digits, special_chars, letters, upper_case, lower_case, spaces, total_chars)
    VALUES (?,?, ?, ?, ?, ?, ?, ?, ?)
    `

	// Execute the query with the values
	_, err := DB.Exec(query, analysis.Vowels, analysis.Consonants, analysis.Digits, analysis.SpecialChars,
		analysis.Letters, analysis.UpperCase, analysis.LowerCase, analysis.Spaces, analysis.TotalChars)
	if err != nil {
		log.Println("Error inserting data into the database:", err)
		return err
	}

	return nil
}

package db

import (
	"log"

	"github.com/malaika-muneer/File-Analyser/models"
)

func (d *dao) InsertAnalysisData(analysis models.FileAnalysis) error {
	log.Printf("Database insert")
	query := `
    INSERT INTO analysis_results (user_id,username,vowels, consonants, digits, special_chars, letters, upper_case, lower_case, spaces, total_chars)
    VALUES (?,?,?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	// Execute the query with the values
	_, err := d.DB.Exec(query, analysis.Id, analysis.Username, analysis.Vowels, analysis.Consonants, analysis.Digits, analysis.SpecialChars,
		analysis.Letters, analysis.UpperCase, analysis.LowerCase, analysis.Spaces, analysis.TotalChars)
	if err != nil {
		log.Println("Error inserting data into the database:", err)
		return err
	}

	return nil
}

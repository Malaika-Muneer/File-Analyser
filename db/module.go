package db

import (
	"database/sql"

	"github.com/malaika-muneer/File-Analyser/models"
)

type DaoLayer interface {
	InsertAnalysisData(analysis models.FileAnalysis) error
	InsertUser(user models.User) error
}

type dao struct {
	DB *sql.DB
}

func NewDao(db *sql.DB) *dao {
	// db := ConnectDB()
	return &dao{
		DB: db,
	}
}

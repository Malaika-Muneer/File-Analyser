package mongodb

import "github.com/malaika-muneer/File-Analyser/models"

type MongoLayer interface {
	InsertAnalysisData(analysis models.FileAnalysis) error
}

type mongoDAO struct {
	Collection any
}

func NewMongoDAO(collection any) *mongoDAO {
	return &mongoDAO{
		Collection: collection,
	}
}

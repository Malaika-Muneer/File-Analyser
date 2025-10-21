package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/malaika-muneer/File-Analyser/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDAO struct {
	Collection *mongo.Collection
}

// Constructor
func NewMongo(collection *mongo.Collection) *MongoDAO {
	return &MongoDAO{Collection: collection}
}

// Insert analysis data into MongoDB
func (m *MongoDAO) InsertAnalysisData(analysis models.FileAnalysis) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := bson.M{
		"_id":           primitive.NewObjectID(),
		"user_id":       analysis.Id,
		"username":      analysis.Username,
		"vowels":        analysis.Vowels,
		"consonants":    analysis.Consonants,
		"digits":        analysis.Digits,
		"special_chars": analysis.SpecialChars,
		"letters":       analysis.Letters,
		"upper_case":    analysis.UpperCase,
		"lower_case":    analysis.LowerCase,
		"spaces":        analysis.Spaces,
		"total_chars":   analysis.TotalChars,
		"chunk_number":  analysis.ChunkNumber,
		"created_at":    time.Now(),
	}

	_, err := m.Collection.InsertOne(ctx, doc)
	if err != nil {
		log.Println("Error inserting analysis data into MongoDB:", err)
		return fmt.Errorf("mongo insert error: %v", err)
	}

	log.Println("Inserted analysis data into MongoDB successfully.")
	return nil
}

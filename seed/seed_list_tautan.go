package seed

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"UNISA_Server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedListTautan(db *mongo.Database) error {
	collection := db.Collection("list_tautan")
	statusCollection := db.Collection("seeding_status")

	// Check if seeding has already been done
	var result bson.M
	err := statusCollection.FindOne(context.Background(), bson.M{"collection": "list_tautan"}).Decode(&result)
	if err == nil {
		// Seeding has already been done for this collection
		return nil
	}

	// Open the JSON file
	file, err := os.Open("data/ListTautan.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the JSON data
	var data []models.ListTautan
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return err
	}

	// Insert the data into the collection
	documents := make([]interface{}, len(data))
	for i, item := range data {
		documents[i] = item
	}

	_, err = collection.InsertMany(context.Background(), documents)
	if err != nil {
		return err
	}

	// Record that seeding has been completed
	_, err = statusCollection.InsertOne(context.Background(), bson.M{
		"collection": "list_tautan",
		"seeded_at":  time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

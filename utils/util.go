package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Struct to hold database and its collections
type DatabaseInfo struct {
	Name        string   `json:"name"`
	Collections []string `json:"collections"`
}


// Function to connect using a MongoDB URI and return databases + collections
func GetDatabasesAndCollections(mongoURI string) ([]DatabaseInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create temporary client (doesn't affect your main DB client)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	dbNames, err := client.ListDatabaseNames(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var result []DatabaseInfo

	for _, dbName := range dbNames {
		db := client.Database(dbName)

		collections, err := db.ListCollectionNames(ctx, map[string]interface{}{})
		if err != nil {
			continue // skip DBs we can't access
		}

		result = append(result, DatabaseInfo{
			Name:        dbName,
			Collections: collections,
		})
	}

	return result, nil
}

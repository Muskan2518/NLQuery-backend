package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := "mongodb+srv://muskan220405:SRxsMqe9UQ9vXVLj@cluster0.lufiaal.mongodb.net/?retryWrites=true&w=majority"

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	log.Println("✅ Connected to MongoDB Atlas")

	Client = client
}

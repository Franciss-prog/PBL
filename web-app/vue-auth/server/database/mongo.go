package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client // Global variable to hold the database client

func ConnectDB() (*mongo.Client, error) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system environment variables")
	}

	// Retrieve MongoDB connection string
	mongoURL := os.Getenv("MONGODB_URL")
	if mongoURL == "" {
		return nil, fmt.Errorf("MONGODB_URL is not set in .env file")
	}

	// Set MongoDB connection options
	clientOptions := options.Client().ApplyURI(mongoURL)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database to check if the connection is successful
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	log.Println("✅ Successfully connected to MongoDB!")

	// Store client in global variable
	DB = client

	return client, nil
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("go-auth").Collection(collectionName)
}

package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var URLCollection *mongo.Collection

func InitMongo() {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	// Ping to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Connected to MongoDB")

	// Select database and collection
	db := client.Database("url_shortener")
	URLCollection = db.Collection("urls")
	Client = client

	// 🔍 Debug print (this is what you asked)
	fmt.Println("📌 Connected DB:", db.Name())
	fmt.Println("📂 Using Collection:", URLCollection.Name())
}

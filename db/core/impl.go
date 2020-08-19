package core

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	client *mongo.Client
}

// Insert creates a new document in DB
func (m *mongoDB) Insert(collection string, doc interface{}) error {
	c := m.client.Database("outdoorsight").Collection(collection)
	if err := c.Insert(doc); err != nil {
		return errors.New(err)
	}
	return nil
}

// ConnectDB : This is helper function to connect mongoDB
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func ConnectDB() *mongo.Client {
	// set client options
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	// connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

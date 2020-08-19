package core

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	outdoorsightDB = "outdoorsight"
)

type MongoDB struct {
	Client *mongo.Client
}

// NewMongoDB creates a new instance of MongoDB structure
func NewMongoDB() MongoDB {
	m := MongoDB{}
	return MongoDB{Client: m.NewClient()}
}

// Insert creates a new document in DB
func (m *MongoDB) Insert(ctx context.Context, collection string, doc interface{}) error {
	c := m.Client.Database(outdoorsightDB).Collection(collection)
	if _, err := c.InsertOne(ctx, doc); err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to insert document in collection %s", collection))
	}
	return nil
}

// Find retrieves the cursor corresponding to the given filter
func (m *MongoDB) Find(ctx context.Context, collection string, filter map[string]interface{}) (*mongo.Cursor, error) {
	c := m.Client.Database(outdoorsightDB).Collection(collection)
	cursor, err := c.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to find document in collection %s", collection))
	}
	defer cursor.Close(ctx)
	return cursor, nil
}

// Delete deletes a document
func (m *MongoDB) Delete(ctx context.Context, collection string, filter map[string]interface{}) error {
	c := m.Client.Database(outdoorsightDB).Collection(collection)
	if _, err := c.DeleteOne(ctx, filter); err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to delete document in collection %s", collection))
	}
	return nil
}

// Update updates an existing document
func (m *MongoDB) Update(ctx context.Context, collection string, filter map[string]interface{}, update interface{}) error {
	c := m.Client.Database(outdoorsightDB).Collection(collection)
	res, err := c.UpdateOne(ctx, filter, update)
	if res.MatchedCount != 1 {
		return errors.Wrap(err, "there is no document corresponding to the filter")
	}
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to update document in collection %s", collection))
	}
	return nil
}

// NewClient creates the connexion to the database and returns a mongo client
func (m *MongoDB) NewClient() *mongo.Client {
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

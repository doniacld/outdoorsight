package core

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	//TODO DONIA pass this URI as an argument
	mongoURI       = "mongodb://172.17.0.2:27017"
	outdoorsightDB = "outdoorsight"
)

// mongoDB holds the mongo client
type mongoDB struct {
	client *mongo.Client
}

// NewDB returns an object implementing DB
func NewDB() (DB, error) {
	m := mongoDB{}
	client, err := m.newClient()
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new mongo client")
	}
	return &mongoDB{client: client}, nil
}

// Insert creates a new document in DB
func (m *mongoDB) Insert(ctx context.Context, collection string, doc interface{}) (interface{}, error) {
	c := m.client.Database(outdoorsightDB).Collection(collection)
	result, err := c.InsertOne(ctx, doc)
	if err != nil {
		return result.InsertedID, errors.Wrap(err, fmt.Sprintf("error while inserting document '%s' in collection '%s'", doc, collection))
	}
	return result.InsertedID, nil
}

// Find retrieves the cursor corresponding to the given filter
func (m *mongoDB) Find(ctx context.Context, collection string, filter map[string]interface{}) (*mongo.Cursor, error) {
	c := m.client.Database(outdoorsightDB).Collection(collection)
	cursor, err := c.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to find document in collection %s", collection))
	}
	return cursor, nil
}

// Update updates an existing document
func (m *mongoDB) Update(ctx context.Context, collection string, filter map[string]interface{}, update bson.D) (int64, int64, error) {
	c := m.client.Database(outdoorsightDB).Collection(collection)
	res, err := c.UpdateOne(ctx, filter, update)
	if err != nil {
		return int64(0), int64(0), errors.Wrap(err, fmt.Sprintf("unable to update document in collection %s", collection))
	}
	return res.MatchedCount, res.ModifiedCount, nil
}

// Delete deletes a document
func (m *mongoDB) Delete(ctx context.Context, collection string, filter map[string]interface{}) (int64, error) {
	c := m.client.Database(outdoorsightDB).Collection(collection)
	res, err := c.DeleteOne(ctx, filter)
	if err != nil {
		return int64(0), errors.Wrap(err, fmt.Sprintf("unable to delete document in collection %s", collection))
	}
	return res.DeletedCount, nil
}

// newClient creates the connection to the database and returns a mongo client
func (m *mongoDB) newClient() (*mongo.Client, error) {
	// set client options
	clientOptions := options.Client().ApplyURI(mongoURI)
	// connect to mongoDB
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to connect to mongo %+v", &clientOptions))
	}

	return client, nil
}

package core

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/doniacld/outdoorsight/errors"

	pkgerrors "github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	//TODO DONIA
	mongoURI       = "mongodb://172.17.0.2:27017"
	outdoorsightDB = "outdoorsight"
)

type mongoDB struct {
	Client *mongo.Client
}

// NewDB creates a new instance of mongoDB structure
func NewDB() DB {
	m := mongoDB{}
	return &mongoDB{Client: m.NewClient()}
}

// Insert creates a new document in DB
func (m *mongoDB) Insert(ctx context.Context, collection string, doc interface{}) *errors.OsError {
	c := m.Client.Database(outdoorsightDB).Collection(collection)
	if _, err := c.InsertOne(ctx, doc); err != nil {
		return errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("unable to insert document in collection %s", collection))
	}
	return nil
}

// Find retrieves the cursor corresponding to the given filter
func (m *mongoDB) Find(ctx context.Context, collection string, filter map[string]interface{}) (*mongo.Cursor, *errors.OsError) {
	c := m.Client.Database(outdoorsightDB).Collection(collection)
	cursor, err := c.Find(ctx, filter)
	if err != nil {
		return nil, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("unable to find document in collection %s", collection))
	}
	return cursor, nil
}

// Update updates an existing document
func (m *mongoDB) Update(ctx context.Context, collection string, filter map[string]interface{}, update bson.D) *errors.OsError {
	c := m.Client.Database(outdoorsightDB).Collection(collection)
	res, err := c.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("unable to update document in collection %s", collection))
	}
	// element not found
	if res.MatchedCount < 1 {
		return errors.New(http.StatusNotFound, fmt.Sprintf("unable to find %s in %s", filter, collection))
	}
	// element is not modified
	if res.ModifiedCount < 1 {
		return errors.New(http.StatusInternalServerError, fmt.Sprintf("error while updating %s in %s", filter, collection))
	}
	return nil
}

// Delete deletes a document
func (m *mongoDB) Delete(ctx context.Context, collection string, filter map[string]interface{}) *errors.OsError {
	c := m.Client.Database(outdoorsightDB).Collection(collection)
	if _, err := c.DeleteOne(ctx, filter); err != nil {
		return errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("unable to delete document in collection %s", collection))
	}
	return nil
}

// NewClient creates the connexion to the database and returns a mongo client
func (m *mongoDB) NewClient() *mongo.Client {
	// set client options
	clientOptions := options.Client().ApplyURI(mongoURI)
	// connect to mongoDB
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(pkgerrors.New(fmt.Sprintf("unable to connect to mongo %q", &clientOptions)))
	}

	return client
}

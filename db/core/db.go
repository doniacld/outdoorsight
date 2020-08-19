package core

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

// DB is the core database interface
type DB interface {
	// DB handlers
	Insert(ctx context.Context, collection string, doc interface{}) error
	Find(ctx context.Context, collection string, filter map[string]interface{}) (*mongo.Cursor, error)
	Delete(ctx context.Context, collection string, filter map[string]interface{}) error
	Update(ctx context.Context, collection string, filter map[string]interface{}, update interface{}) error

	// Mongo connexion
	NewClient() *mongo.Client
}

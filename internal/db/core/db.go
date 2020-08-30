package core

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DB is the core database interface
type DB interface {
	Insert(ctx context.Context, collection string, doc interface{}) (interface{}, error)
	Find(ctx context.Context, collection string, filter map[string]interface{}) (*mongo.Cursor, error)
	Update(ctx context.Context, collection string, filter map[string]interface{}, update bson.D) (int64, int64, error)
	Delete(ctx context.Context, collection string, filter map[string]interface{}) (int64, error)
}

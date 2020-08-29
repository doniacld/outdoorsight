package core

import (
	"context"

	"github.com/doniacld/outdoorsight/internal/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DB is the core database interface
type DB interface {
	Insert(ctx context.Context, collection string, doc interface{}) *errors.OsError
	Find(ctx context.Context, collection string, filter map[string]interface{}) (*mongo.Cursor, *errors.OsError)
	Update(ctx context.Context, collection string, filter map[string]interface{}, update bson.D) *errors.OsError
	Delete(ctx context.Context, collection string, filter map[string]interface{}) *errors.OsError
}

package core

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// DB is the core database interface
type DB interface {
	Insert(collection string, doc interface{}) error
	ConnectDB() *mongo.Client
}




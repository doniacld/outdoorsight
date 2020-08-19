package db

import (
	"context"
	"github.com/doniacld/outdoorsight/db/core"
)

// DB is the database interface
// All the access to the database are hidden here
type DB interface {
	// Manage spot
	AddSpot(ctx context.Context, details SpotDetails) error
	GetSpot(ctx context.Context, spot string) (SpotDetails, error)
	DeleteSpot(ctx context.Context, spot string) error
	UpdateSpot(ctx context.Context, details SpotDetails) error
}

// OutdoorsightDB is the structure holding the core DB
type OutdoorsightDB struct {
	DB core.MongoDB
}

// NewClient creates the OutdoorsightDB structure
func New() OutdoorsightDB {
	mongoDB := core.NewMongoDB()
	return OutdoorsightDB{
		mongoDB,
	}
}

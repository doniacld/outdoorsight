package db

import (
	"context"
	"github.com/doniacld/outdoorsight/internal/db/core"
	"github.com/doniacld/outdoorsight/internal/spot"
)

// DB is the database interface
// All the access to the database are hidden here
type DB interface {
	// Manage spot
	AddSpot(ctx context.Context, details SpotDetails) error
	GetSpot(ctx context.Context, spotName string) (spot.Details, error)
	DeleteSpot(ctx context.Context, spotName string) error
	UpdateSpot(ctx context.Context, spotName string, update SpotDetails) error
}

// OutdoorsightDB is the structure holding the core DB
type OutdoorsightDB struct {
	core.DB
}

// NewClient creates the OutdoorsightDB structure
func New() DB {
	mongoDB := core.NewDB()
	osDB := OutdoorsightDB{mongoDB}
	return &osDB
}

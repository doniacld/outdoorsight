package db

import (
	"context"

	"github.com/doniacld/outdoorsight/internal/db/core"
	"github.com/doniacld/outdoorsight/internal/errors"
	"github.com/doniacld/outdoorsight/internal/spot"
)

// DB is the database interface
// All the access to the database are hidden here
type DB interface {
	// Manage spot
	AddSpot(ctx context.Context, details SpotDetails) *errors.OsError
	GetSpot(ctx context.Context, spotName string) (spot.Details, *errors.OsError)
	DeleteSpot(ctx context.Context, spotName string) *errors.OsError
	UpdateSpot(ctx context.Context, spotName string, update SpotDetails) *errors.OsError
}

// OutdoorsightDB is the structure holding the core DB
type OutdoorsightDB struct {
	core.CoreDB
}

// NewClient creates the OutdoorsightDB structure
func New() DB {
	coreDB := core.NewDB()
	osDB := OutdoorsightDB{coreDB}
	return &osDB
}

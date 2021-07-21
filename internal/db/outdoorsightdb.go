package db

import (
	"context"
	"log"

	"github.com/doniacld/outdoorsight/internal/db/core"
	"github.com/doniacld/outdoorsight/internal/spot"

	"github.com/pkg/errors"
)

// DB is the database interface
// All the accesses to the database are hidden here
type DB interface {
	// Manage spot
	AddSpot(ctx context.Context, details SpotDetails) (interface{}, error)
	GetSpot(ctx context.Context, spotName string) (*spot.Details, error)
	UpdateSpot(ctx context.Context, spotName string, update SpotDetails) (int64, int64, error)
	DeleteSpot(ctx context.Context, spotName string) (int64, error)
}

// OutdoorsightDB is the structure holding the core DB
type OutdoorsightDB struct {
	core.DB
}

// New creates the DB interface
func New() (DB, error) {
	coreDB, err := core.NewDB()
	if err != nil {
		return nil, errors.Wrap(err, "error while creating a new instance of DB")
	}
	osDB := OutdoorsightDB{coreDB}
	log.Print("create a new instance of DB")
	return &osDB, nil
}

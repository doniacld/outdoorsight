package db

import (
	"context"

	"github.com/doniacld/outdoorsight/db/core"
)

// DB is the DB interface
type DB interface {
	// Manage spot
	AddSpot(ctx context.Context, details SpotDetails) error
	GetSpot(ctx context.Context, spot string) (SpotDetails, error)
	DeleteSpot(ctx context.Context, spot string) error
}

// outdoorsightDB is the strcture holding the core DB
type outdoorsightDB struct {
	core.DB
}

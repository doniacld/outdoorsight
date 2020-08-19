package db

import (
	"context"
	"errors"
)

const (
	spotsCollection = "spots"
)

// AddSpot adds a spot with all its details in database
func (db *outdoorsightDB) AddSpot(_ context.Context, details SpotDetails) error {
	if err := db.Insert(spotsCollection, details); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// GetSpot retrieves a given spot with its details from databas
func (db *outdoorsightDB) GetSpot(ctx context.Context, spot string) (SpotDetails, error) {
	return SpotDetails{}, nil
}

// DeleteSpot deletes a spot from
func (db *outdoorsightDB) DeleteSpot(ctx context.Context, spot string) error {
	return nil
}

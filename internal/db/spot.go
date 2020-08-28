package db

import (
	"context"
	"fmt"
	"github.com/doniacld/outdoorsight/internal/spot"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/pkg/errors"
)

const (
	spotsCollection = "spots"
)

// AddSpot adds a spot with all its details in database
func (os *OutdoorsightDB) AddSpot(ctx context.Context, details SpotDetails) error {

	if err := os.Insert(ctx, spotsCollection, details); err != nil {
		return errors.Wrap(err, "unable to add spot in DB")
	}
	return nil
}

// GetSpot retrieves a given spot with its details from database
func (os *OutdoorsightDB) GetSpot(ctx context.Context, spotName string) (spot.Details, error) {
	// retrieve spot details in DB
	cursor, err := os.Find(ctx, spotsCollection, spotNameFilter(spotName))
	if err != nil {
		return spot.Details{}, errors.Wrap(err, fmt.Sprintf("unable to get spot %s", spotName))
	}

	// convert the result which is a cursor in a spot.Details structure
	var spotDetails spot.Details
	for cursor.Next(ctx) {
		err := cursor.Decode(&spotDetails)
		if err != nil {
			return spot.Details{}, errors.Wrap(err, "unable to decode cursor")
		}
	}
	// case the name is not found
	if len(spotDetails.Name) == 0 {
		return spot.Details{}, errors.New(fmt.Sprintf("spotDetails '%s' is not found", spotName))

	}
	return spotDetails, nil
}

// UpdateSpot updates a spot in database
func (os *OutdoorsightDB) UpdateSpot(ctx context.Context, spotName string, sd SpotDetails) error {
	update := bson.D{{"$set", sd}}
	if err := os.Update(ctx, spotsCollection, spotNameFilter(spotName), update); err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to update spot %s in DB", spotName))
	}
	return nil
}

// DeleteSpot deletes a spot from database
func (os *OutdoorsightDB) DeleteSpot(ctx context.Context, spotName string) error {
	if err := os.Delete(ctx, spotsCollection, spotNameFilter(spotName)); err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to delete spot %s from DB", spotName))
	}
	return nil
}

// spotNameFilter creates a filter on a spot name
func spotNameFilter(spotName string) map[string]interface{} {
	return map[string]interface{}{
		"name": spotName,
	}
}

// TODO DONIA try again to use it ?
// decodeCursor transforms a mongo cursor into the interface that you want
func decodeCursor(ctx context.Context, cursor *mongo.Cursor, format interface{}) (interface{}, error) {
	for cursor.Next(ctx) {
		err := cursor.Decode(&format)
		if err != nil {
			return nil, errors.Wrap(err, "unable to decode cursor")
		}
	}
	return format, nil
}

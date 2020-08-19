package db

import (
	"context"
	"fmt"
	"github.com/doniacld/outdoorsight/spot"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	spotsCollection = "spots"
)

// AddSpot adds a spot with all its details in database
func (os *OutdoorsightDB) AddSpot(ctx context.Context, details SpotDetails) error {
	if err := os.DB.Insert(ctx, spotsCollection, details); err != nil {
		return errors.Wrap(err, "unable to add spot in DB")
	}
	return nil
}

// GetSpot retrieves a given spot with its details from database
func (os *OutdoorsightDB) GetSpot(ctx context.Context, spotName string) (spot.Details, error) {
	filter := bson.M{"name": spotName}
	cursor, err := os.DB.Find(ctx, spotsCollection, filter)
	if err != nil {
		return spot.Details{}, errors.Wrap(err, fmt.Sprintf("unable to get spot %s", spotName))
	}

	// convert the result which is a cursor in a spotDetails structure
	var spotDetails spot.Details
	for cursor.Next(ctx) {
		err := cursor.Decode(&spotDetails)
		if err != nil {
			return spot.Details{}, errors.Wrap(err, fmt.Sprintf("unable to decode spot %s from DB", spotName))
		}
	}

	return spotDetails, nil
}

// DeleteSpot deletes a spot from database
func (os *OutdoorsightDB) DeleteSpot(ctx context.Context, spotName string) error {
	filter := bson.M{"name": spotName}
	if err := os.DB.Delete(ctx, spotsCollection, filter); err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to delete spot %s from DB", spotName))
	}
	return nil
}

// UpdateSpot updates a spot in database
func (os *OutdoorsightDB) UpdateSpot(ctx context.Context, spotName string, details SpotDetails) error {
	filter := bson.M{"name": spotName}
	// TODO DONIA update all the given fields
	update := bson.D{{"$set", bson.D{{"name", details.Name}}}}

	if err := os.DB.Update(ctx, spotsCollection, filter, update); err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to update spot %s in DB", details.Name))
	}
	return nil
}

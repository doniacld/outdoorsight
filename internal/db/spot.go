package db

import (
	"context"
	"fmt"
	"github.com/doniacld/outdoorsight/internal/spot"
	"log"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	spotsCollection = "spots"
)

// AddSpot adds a spot with all its details in database
func (ods *OutdoorsightDB) AddSpot(ctx context.Context, details SpotDetails) (interface{}, error) {
	res, err := ods.Insert(ctx, spotsCollection, details)
	if err != nil {
		return nil, errors.Wrap(err, "unable to add spot in DB")
	}
	return res, nil
}

// GetSpot retrieves a given spot with its details from database
func (ods *OutdoorsightDB) GetSpot(ctx context.Context, spotName string) (*spot.Details, error) {
	// retrieve spot details in DB
	cursor, err := ods.Find(ctx, spotsCollection, spotNameFilter(spotName))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to get spot %s", spotName))
	}

	// check if cursor is empty, if yes returns an empty spot.Details
	if cursor.Next(ctx) == false {
		log.Print("enter is cursor next is false")
		return nil, nil
	}
	log.Printf("retrieved cursor from database '%v'", cursor)
	// convert the result which is a cursor into a spot.Details structure
	spotDetails := spot.Details{}
	if err := cursor.Decode(&spotDetails); err != nil {
		return nil, errors.Wrap(err, "unable to decode cursor")
	}
	log.Printf("retrieved spotDetails from database '%v'", &spotDetails)
	return &spotDetails, nil
}

// UpdateSpot updates a spot in database
func (ods *OutdoorsightDB) UpdateSpot(ctx context.Context, spotName string, sd SpotDetails) (int64, int64, error) {
	fmt.Println("spot details", sd)
	update := bson.D{{"$set", sd}}
	matchedCount, modifiedCount, err := ods.Update(ctx, spotsCollection, spotNameFilter(spotName), update)
	if err != nil {
		return 0, 0, errors.Wrap(err, fmt.Sprintf("unable to update spot %s in DB", spotName))
	}
	return matchedCount, modifiedCount, nil
}

// DeleteSpot deletes a spot from database
func (ods *OutdoorsightDB) DeleteSpot(ctx context.Context, spotName string) (int64, error) {
	deletedCount, err := ods.Delete(ctx, spotsCollection, spotNameFilter(spotName))
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("unable to delete spot %s from DB", spotName))
	}
	return deletedCount, nil
}

// filter represents the applied filter passed to the DB
type filter map[string]interface{}

// spotNameFilter creates a filter on a spot name
func spotNameFilter(spotName string) filter {
	return filter{"name": spotName}
}

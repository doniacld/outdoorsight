package addspot

import (
	"context"
	"encoding/json"
	"github.com/doniacld/outdoorsight/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"github.com/pkg/errors"
)

// AddSpot adds a spot to your list
func AddSpot(request AddSpotRequest) (AddSpotResponse, error) {
	spotDetailsDB, err := convertToSpotDetailsDB(request)
	if err != nil {
		return AddSpotResponse{}, errors.Wrap(err, "error while converting to spotDetails DB structure")
	}

	// call the database to add the details
	osDB := db.New()
	if err := osDB.AddSpot(context.TODO(), spotDetailsDB); err != nil {
		panic(errors.Wrap(err, "unable to add spot"))
	}
	// call the database to get the details (a way to verify that we really added the data)
	spotDetails, err := osDB.GetSpot(context.TODO(), spotDetailsDB.Name)
	if err != nil {
		return AddSpotResponse{}, errors.Wrapf(err, "unable to get spot %s details", spotDetails.Name)
	}
	response := AddSpotResponse(spotDetails)
	return response, nil
}

// convertToSpotDetailsDB converts the request to spotDetails DB structure
func convertToSpotDetailsDB(request AddSpotRequest) (db.SpotDetails, error) {
	data, err := json.Marshal(&request)
	if err != nil {
		return db.SpotDetails{}, errors.Wrapf(err, "error while marshalling request '%s'", request)
	}

	var spotDetailsDB db.SpotDetails
	if err := bson.UnmarshalExtJSON(data, true, &spotDetailsDB); err != nil {
		// if err := bson.Unmarshal(data, &spotDetailsDB); err != nil {
		return db.SpotDetails{}, errors.Wrapf(err, "error while unmarshalling spotDetailsDB from request '%s'", request)
	}
	log.Printf("convert addSpotRequest to spotDetailsDB: %q, %T", spotDetailsDB, spotDetailsDB)
	return spotDetailsDB, nil
}

package addspot

import (
	"context"
	"encoding/json"

	"github.com/doniacld/outdoorsight/internal/db"

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
	if err := osDB.AddSpot(context.TODO(), db.SpotDetails{}); err != nil {
		panic(errors.Wrap(err, "unable to add spot"))
	}
	// call the database to get the details (a way to verify that we really added the data)
	spotDetails, err := osDB.GetSpot(context.TODO(), spotDetailsDB.Name)
	if err != nil {
		return AddSpotResponse{}, errors.Wrapf(err, "unable to get spot %s", spotDetails.Name)
	}

	response, err := convertToAddSpotResponse(spotDetailsDB)
	if err != nil {
		return AddSpotResponse{}, errors.Wrapf(err, "error while converting to %s", spotDetails.Name)
	}

	return response, nil
}

// convertToSpotDetailsDB converts the request to spotDetails DB structure
func convertToSpotDetailsDB(request AddSpotRequest) (db.SpotDetails, error) {
	data, err := json.Marshal(&request)
	if err != nil {
		return db.SpotDetails{}, errors.Wrapf(err, "error while marshalling request '%s'", request)
	}

	var spotDetailsDB db.SpotDetails
	if err := json.Unmarshal(data, &spotDetailsDB); err != nil {
		return db.SpotDetails{}, errors.Wrapf(err, "error while unmarshalling spotDetailsDB from request '%s'", request)
	}
	return spotDetailsDB, nil
}

// convertToAddSpotResponse converts the spotDetails DB to AddSpotResponse
func convertToAddSpotResponse(spotDetailsDB db.SpotDetails) (AddSpotResponse, error) {
	data, err := json.Marshal(&spotDetailsDB)
	if err != nil {
		return AddSpotResponse{}, errors.Wrapf(err, "error while marshalling spotDetailsDB '%s'", spotDetailsDB)
	}

	var response AddSpotResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return AddSpotResponse{}, errors.Wrapf(err, "error while unmarshalling request from '%s'", spotDetailsDB)
	}
	return response, nil
}

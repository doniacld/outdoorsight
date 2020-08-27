package updatespot

import (
	"context"
	"encoding/json"

	"github.com/doniacld/outdoorsight/internal/db"

	"github.com/pkg/errors"
)

// UpdateSpot returns all the details about a given spot
func UpdateSpot(ctx context.Context, request UpdateSpotRequest) (UpdateSpotResponse, error) {
	// convert the request into spotDetails DB structure
	spotDetailsDB, err := convertToSpotDetailsDB(request)
	if err != nil {
		return UpdateSpotResponse{}, errors.Wrapf(err, "unable to delete spot %s", request.Name)
	}

	// update the spot in DB
	osDB := db.New()
	if err := osDB.UpdateSpot(ctx, request.Name, spotDetailsDB); err != nil {
		return UpdateSpotResponse{}, errors.Wrapf(err, "unable to update spot %s", request.Name)
	}

	// get the updated spot name
	if request.Name != spotDetailsDB.Name {
		request.Name = spotDetailsDB.Name
	}
	spotDetails, err := osDB.GetSpot(ctx, request.Name)
	if err != nil {
		return UpdateSpotResponse{}, errors.Wrapf(err, "unable to get spot %s", request.Name)
	}

	response := UpdateSpotResponse(spotDetails)
	return response, nil
}

// convertToSpotDetailsDB converts the request to spotDetails DB structure
func convertToSpotDetailsDB(request UpdateSpotRequest) (db.SpotDetails, error) {
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

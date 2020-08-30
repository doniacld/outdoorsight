package updatespot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/doniacld/outdoorsight/internal/db"
	"github.com/doniacld/outdoorsight/internal/errors"
)

// UpdateSpot returns all the details about a given spot
func UpdateSpot(ctx context.Context, request UpdateSpotRequest) (UpdateSpotResponse, *errors.OsError) {

	osDB := db.New()

	// call the database to get the details (a way to verify that we really added the data)
	spotDetails, err := osDB.GetSpot(ctx, request.Name)
	if err != nil {
		return UpdateSpotResponse{}, errors.Wrap(err, fmt.Sprintf("unable to get spot '%s' details", request.Name))
	}
	if spotDetails.Name == "" {
		return UpdateSpotResponse{}, errors.New(http.StatusNotFound, fmt.Sprintf("spot '%s' not found", request.Name))
	}

	// convert the request into spotDetails DB structure
	spotDetailsDB, err := convertToSpotDetailsDB(request)
	if err != nil {
		return UpdateSpotResponse{}, errors.Wrap(err, fmt.Sprintf("unable to update spot %s", request.Name))
	}

	// update the spot in DB
	if err := osDB.UpdateSpot(ctx, request.Name, spotDetailsDB); err != nil {
		return UpdateSpotResponse{}, errors.Wrap(err, fmt.Sprintf("unable to update spot %s", request.Name))
	}

	// get the updated spot name
	if request.Name != spotDetailsDB.Name {
		request.Name = spotDetailsDB.Name
	}
	spotDetails, err = osDB.GetSpot(ctx, request.Name)
	if err != nil {
		return UpdateSpotResponse{}, errors.Wrap(err, fmt.Sprintf("unable to get spot %s", request.Name))
	}

	response := UpdateSpotResponse(spotDetails)
	return response, nil
}

// convertToSpotDetailsDB converts the request to spotDetails DB structure
func convertToSpotDetailsDB(request UpdateSpotRequest) (db.SpotDetails, *errors.OsError) {
	data, err := json.Marshal(&request)
	if err != nil {
		return db.SpotDetails{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("error while marshalling request '%q'", request))
	}

	var spotDetailsDB db.SpotDetails
	if err := json.Unmarshal(data, &spotDetailsDB); err != nil {
		return db.SpotDetails{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("error while unmarshalling spotDetailsDB from request '%q'", request))
	}
	return spotDetailsDB, nil
}

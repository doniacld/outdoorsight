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
func UpdateSpot(ctx context.Context, request UpdateSpotRequest, odsDB db.DB) (UpdateSpotResponse, *errors.ODSError) {
	// call the database to get the details (a way to verify that we really added the data)
	spotDetails, err := odsDB.GetSpot(ctx, request.Name)
	if err != nil {
		return UpdateSpotResponse{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("error whild calling getSpot DB for spot '%s'", request.Name))
	}

	// if the spotDetails are empty, it means the endpoint does not exist
	if spotDetails == nil {
		return UpdateSpotResponse{}, errors.New(http.StatusNotFound, fmt.Sprintf("spot '%s' does not exist", request.Name))
	}

	// convert the request into spotDetails DB structure
	spotDetailsDB, er := convertToSpotDetailsDB(request)
	if er != nil {
		return UpdateSpotResponse{}, er.Wrap(fmt.Sprintf("unable to update spot %s", request.Name))
	}

	// update the spot in DB
	if _, _, err := odsDB.UpdateSpot(ctx, request.Name, spotDetailsDB); err != nil {
		return UpdateSpotResponse{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("unable to update spot %s", request.Name))
	}

	// retrieve the updated details
	spotDetails, err = odsDB.GetSpot(ctx, request.Name)
	if err != nil {
		return UpdateSpotResponse{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("unable to get spot '%s' after insertion", request.Name))
	}
	sd := *spotDetails
	// convert the details into an update spot response
	response := UpdateSpotResponse{Name: sd.Name, Routes: sd.Routes, Metadata: sd.Metadata}

	return response, nil
}

// convertToSpotDetailsDB converts the request to spotDetails DB structure
func convertToSpotDetailsDB(request UpdateSpotRequest) (db.SpotDetails, *errors.ODSError) {
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

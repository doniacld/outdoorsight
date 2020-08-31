package addspot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/doniacld/outdoorsight/internal/db"
	"github.com/doniacld/outdoorsight/internal/errors"

	"go.mongodb.org/mongo-driver/bson"
)

// AddSpot adds a spot to your list
func AddSpot(ctx context.Context, request AddSpotRequest, odsDB db.DB) (AddSpotResponse, *errors.ODSError) {
	spotDetailsDB, err := convertToSpotDetailsDB(request)
	if err != nil {
		return AddSpotResponse{}, err.Wrap("error while converting to spotDetails DB structure")
	}

	// call the database to get the details (a way to verify that we really added the data)
	spotDetails, er := odsDB.GetSpot(ctx, spotDetailsDB.Name)
	if er != nil {
		return AddSpotResponse{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("unable to check for an already existing spot details with name '%s'", request.Name))
	}

	// if spotDetails is not empty it means the spot already exists in database
	if spotDetails != nil {
		log.Printf("retrieved spotDetails not empty %q", spotDetails)
		return AddSpotResponse{}, errors.New(http.StatusNotFound, fmt.Sprintf("spot '%s' already exists in database", spotDetails.Name))
	}
	// insert spot in database
	if _, err := odsDB.AddSpot(ctx, spotDetailsDB); err != nil {
		return AddSpotResponse{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("unable to add spot %s", request.Name))
	}
	log.Printf("spot '%s' is inserted in DB", request.Name)

	// call the database to get the details (a way to verify that we really added the data)
	spotDetails, er = odsDB.GetSpot(ctx, request.Name)
	if er != nil {
		return AddSpotResponse{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("unable to get spot %s details", request.Name))
	}
	log.Printf("retrieved '%q' &&& '%p'", spotDetails, &spotDetails)

	// convert DB response into addSpotResponse structure
	response := AddSpotResponse(*spotDetails)
	log.Printf("spotDetails convert to AddSpotResponse '%p', '%q'", &spotDetails, AddSpotResponse(*spotDetails))
	return response, nil
}

// convertToSpotDetailsDB converts the request to spotDetails DB structure
func convertToSpotDetailsDB(request AddSpotRequest) (db.SpotDetails, *errors.ODSError) {
	data, err := json.Marshal(&request)
	if err != nil {
		return db.SpotDetails{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("error while marshalling request '%q'", request))
	}

	var spotDetailsDB db.SpotDetails
	if err := bson.UnmarshalExtJSON(data, true, &spotDetailsDB); err != nil {
		// if err := bson.Unmarshal(data, &spotDetailsDB); err != nil {
		return db.SpotDetails{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("error while unmarshalling spotDetailsDB from request '%q'", request))
	}
	log.Printf("addSpot request is converted to spotDetails db format for '%s'", request.Name)
	return spotDetailsDB, nil
}

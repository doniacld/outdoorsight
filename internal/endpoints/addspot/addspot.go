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
func AddSpot(request AddSpotRequest) (AddSpotResponse, *errors.OsError) {
	spotDetailsDB, err := convertToSpotDetailsDB(request)
	if err != nil {
		return AddSpotResponse{}, errors.Wrap(err, "error while converting to spotDetails DB structure")
	}

	// call the database to add the details
	osDB := db.New()
	if err := osDB.AddSpot(context.TODO(), spotDetailsDB); err != nil {
		return AddSpotResponse{}, errors.Wrap(err, fmt.Sprintf("unable to add spot %s", request.Name))
	}
	// call the database to get the details (a way to verify that we really added the data)
	spotDetails, err := osDB.GetSpot(context.TODO(), spotDetailsDB.Name)
	if err != nil {
		return AddSpotResponse{}, errors.Wrap(err, fmt.Sprintf("unable to get spot %s details", spotDetails.Name))
	}

	// convert DB response into addSpotResponse structure
	response := AddSpotResponse(spotDetails)
	return response, nil
}

// convertToSpotDetailsDB converts the request to spotDetails DB structure
func convertToSpotDetailsDB(request AddSpotRequest) (db.SpotDetails, *errors.OsError) {
	data, err := json.Marshal(&request)
	if err != nil {
		return db.SpotDetails{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("error while marshalling request '%q'", request))
	}

	var spotDetailsDB db.SpotDetails
	if err := bson.UnmarshalExtJSON(data, true, &spotDetailsDB); err != nil {
		// if err := bson.Unmarshal(data, &spotDetailsDB); err != nil {
		return db.SpotDetails{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("error while unmarshalling spotDetailsDB from request '%q'", request))
	}
	log.Printf("convert addSpotRequest to spotDetailsDB: %q, %T", spotDetailsDB, spotDetailsDB)
	return spotDetailsDB, nil
}

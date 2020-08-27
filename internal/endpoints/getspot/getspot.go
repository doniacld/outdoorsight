package getspot

import (
	"context"

	"github.com/pkg/errors"

	"github.com/doniacld/outdoorsight/internal/db"
)

// GetSpot returns all the details on a given spot
func GetSpot(request GetSpotRequest) (GetSpotResponse, error) {

	// call the database
	osDB := db.New()
	spotDetails, err := osDB.GetSpot(context.TODO(), request.SpotName)
	if err != nil {
		return GetSpotResponse{}, errors.Wrapf(err, "unable to get spot %s", request.SpotName)
	}
	// convert the spot.Details to the response
	response := GetSpotResponse(spotDetails)
	return response, nil
}

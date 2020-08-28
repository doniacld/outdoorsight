package getspot

import (
	"context"
	"fmt"
	"github.com/doniacld/outdoorsight/errors"

	"github.com/doniacld/outdoorsight/internal/db"
)

// GetSpot returns all the details on a given spot
func GetSpot(request GetSpotRequest) (GetSpotResponse, *errors.OsError) {
	// call the method database
	osDB := db.New()
	spotDetails, err := osDB.GetSpot(context.TODO(), request.SpotName)
	if err != nil {
		return GetSpotResponse{}, errors.Wrap(err, fmt.Sprintf("unable to get spot %s", request.SpotName))
	}

	// convert the spot.Details to the response
	response := GetSpotResponse(spotDetails)
	return response, nil
}

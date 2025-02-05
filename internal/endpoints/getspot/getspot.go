package getspot

import (
	"context"
	"fmt"
	"github.com/doniacld/outdoorsight/internal/errors"
	"log"
	"net/http"

	"github.com/doniacld/outdoorsight/internal/db"
)

// GetSpot returns all the details on a given spot
func GetSpot(ctx context.Context, request GetSpotRequest, odsDB db.DB) (GetSpotResponse, *errors.ODSError) {
	// call the method database
	spotDetails, err := odsDB.GetSpot(ctx, request.SpotName)
	if err != nil {
		return GetSpotResponse{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("error while calling getSpot DB for spot %s", request.SpotName))
	}
	// GetSpot from DB return an empty pointer if the spot is not found
	if spotDetails == nil {
		log.Printf("spot '%s' is not found", request.SpotName)
		return GetSpotResponse{}, errors.New(http.StatusNotFound, fmt.Sprintf("spot '%s' does not exist", request.SpotName))
	}
	log.Printf("retrieved spotDetails for '%s'", request.SpotName)

	// convert the spot.Details to the response
	response := GetSpotResponse(*spotDetails)
	return response, nil
}

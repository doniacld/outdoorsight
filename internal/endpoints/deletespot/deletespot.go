package deletespot

import (
	"context"
	"fmt"
	"net/http"

	"github.com/doniacld/outdoorsight/internal/db"
	"github.com/doniacld/outdoorsight/internal/errors"
)

// DeleteSpot deletes a given spot
// endpoint is idempotent which means even the image is not deleted, we return the same response
func DeleteSpot(ctx context.Context, request DeleteSpotRequest, odsDB db.DB) (DeleteSpotResponse, *errors.ODSError) {
	_, err := odsDB.DeleteSpot(ctx, request.SpotName)
	if err != nil {
		return DeleteSpotResponse{}, errors.NewFromError(http.StatusInternalServerError, err, fmt.Sprintf("unable to delete spot %s", request.SpotName))
	}
	return DeleteSpotResponse{}, nil
}

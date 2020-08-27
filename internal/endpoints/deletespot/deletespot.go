package deletespot

import (
	"context"
	"github.com/pkg/errors"

	"github.com/doniacld/outdoorsight/internal/db"
)

// DeleteSpot deletes a given spot

func DeleteSpot(ctx context.Context, request DeleteSpotRequest) (DeleteSpotResponse, error) {
	osDB := db.New()
	err := osDB.DeleteSpot(ctx, request.SpotName)
	if err != nil {
		return DeleteSpotResponse{}, errors.Wrapf(err, "unable to delete spot %s", request.SpotName)
	}
	return DeleteSpotResponse{}, nil
}
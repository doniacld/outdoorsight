package deletespot

import (
	"context"
	"fmt"

	"github.com/doniacld/outdoorsight/internal/db"
	"github.com/doniacld/outdoorsight/internal/errors"
)

// DeleteSpot deletes a given spot

func DeleteSpot(ctx context.Context, request DeleteSpotRequest) (DeleteSpotResponse, *errors.OsError) {
	osDB := db.New()
	err := osDB.DeleteSpot(ctx, request.SpotName)
	if err != nil {
		return DeleteSpotResponse{}, errors.Wrap(err, fmt.Sprintf("unable to delete spot %s", request.SpotName))
	}
	return DeleteSpotResponse{}, nil
}

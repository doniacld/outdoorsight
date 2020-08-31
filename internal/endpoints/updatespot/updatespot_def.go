package updatespot

import (
	"github.com/doniacld/outdoorsight/internal/errors"
	"net/http"

	"github.com/doniacld/outdoorsight/internal/endpointdef"
	"github.com/doniacld/outdoorsight/internal/endpoints"
	"github.com/doniacld/outdoorsight/internal/spot"
)

// UpdateSpotMeta holds the endpoint information
var UpdateSpotMeta = endpointdef.New(
	"UpdateSpot",
	"/spots/{"+endpoints.ParamSpotName+"}",
	http.MethodPut,
	http.StatusOK,
)

// UpdateSpotRequest is the request
type UpdateSpotRequest spot.Details

// UpdateSpotResponse is the response
type UpdateSpotResponse spot.Details

func (request UpdateSpotRequest) Validate() *errors.ODSError {
	r := spot.Details(request)
	if err := r.Validate(); err != nil {
		return errors.NewFromError(http.StatusBadRequest, err, "error while validating updateSpot request")
	}
	return nil
}

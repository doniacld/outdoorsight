package updatespot

import (
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
var UpdateSpotRequest spot.Details

// UpdateSpotResponse is the response
var UpdateSpotResponse struct{}

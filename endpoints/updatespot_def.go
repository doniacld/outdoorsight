package endpoints

import (
	"net/http"

	"github.com/doniacld/outdoorsight/endpointdef"
	"github.com/doniacld/outdoorsight/spot"
)

// UpdateSpotMeta holds the endpoint information
var UpdateSpotMeta = endpointdef.New(
	"UpdateSpot",
	"/spots/{"+ParamSpotName+"}",
	http.MethodPost,
	http.StatusOK,
)

// UpdateSpotRequest is the request
var UpdateSpotRequest spot.Details

// UpdateSpotResponse is the response
var UpdateSpotResponse struct{}

package endpoints

import (
	"net/http"

	"github.com/doniacld/outdoorsight/endpointdef"
	"github.com/doniacld/outdoorsight/spot"
)

// AddSpotMeta holds the endpoint information
var AddSpotMeta = endpointdef.New(
	"addSpot",
	"/spots",
	http.MethodPut,
	http.StatusCreated,
)

// AddSpotRequest is the request structure
var AddSpotRequest spot.Details

// AddSpotResponse holds the response structure
var AddSpotResponse struct{}

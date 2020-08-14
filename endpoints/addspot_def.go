package endpoints

import (
	"net/http"

	"github.com/doniacld/outdoorsight/endpointdef"
	"github.com/doniacld/outdoorsight/spot"
)

// AddSpotMeta hold the endpoint information
var AddSpotMeta = endpointdef.New(
	"addSpot",
	"/spots",
	http.MethodPost,
	http.StatusCreated,
)

// AddSpotRequest is the request structure
var AddSpotRequest spot.SpotDetails

// AddSpotResponse holds the response structure
var AddSpotResponse struct{}

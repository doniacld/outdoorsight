package endpoints

import (
	"net/http"

	"github.com/doniacld/outdoorsight/endpointdef"
	"github.com/doniacld/outdoorsight/spot"
)

// DeleteSpotMeta hold the endpoint information
var DeleteSpotMeta = endpointdef.New(
	"DeleteSpotDetails",
	"/spots/{"+ParamSpotName+"}",
	http.MethodDelete,
	http.StatusNoContent,
)

// DeleteSpotRequest is the request structure
var DeleteSpotRequest struct {
	SpotName string `json:"spotName"`
}

// DeleteSpotResponse holds the response structure
var DeleteSpotResponse spot.Details

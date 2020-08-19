package endpoints

import (
	"net/http"

	"github.com/doniacld/outdoorsight/endpointdef"
	"github.com/doniacld/outdoorsight/spot"
)

const (
	ParamSpotName = "spotName"
)

// GetSpotMeta hold the endpoint information
var GetSpotMeta = endpointdef.New(
	"getSpotDetails",
	"/spots/{"+ParamSpotName+"}",
	http.MethodGet,
	http.StatusOK,
)

// GetSpotRequest is the request structure
var GetSpotRequest struct {
	SpotName string `json:"spotName"`
}

// GetSpotResponse holds the response structure
var GetSpotResponse spot.Details

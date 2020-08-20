package getspot

import (
	"net/http"

	"github.com/doniacld/outdoorsight/internal/endpointdef"
	"github.com/doniacld/outdoorsight/internal/endpoints"
	"github.com/doniacld/outdoorsight/internal/spot"
)

// GetSpotMeta holds the endpoint information
var GetSpotMeta = endpointdef.New(
	"getSpotDetails",
	"/spots/{"+endpoints.ParamSpotName+"}",
	http.MethodGet,
	http.StatusOK,
)

// GetSpotRequest is the request structure
var GetSpotRequest struct {
	SpotName string `json:"spotName"`
}

// GetSpotResponse holds the response structure
var GetSpotResponse spot.Details

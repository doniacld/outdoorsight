package deletespot

import (
	"net/http"

	"github.com/doniacld/outdoorsight/internal/endpointdef"
	"github.com/doniacld/outdoorsight/internal/endpoints"
)

// DeleteSpotMeta holds the endpoint information
var DeleteSpotMeta = endpointdef.New(
	"DeleteSpotDetails",
	"/spots/{"+endpoints.ParamSpotName+"}",
	http.MethodDelete,
	http.StatusNoContent,
)

// DeleteSpotRequest is the request structure
type DeleteSpotRequest struct {
	SpotName string `json:"spotName"`
}

// DeleteSpotResponse holds the response structure
type DeleteSpotResponse struct {}

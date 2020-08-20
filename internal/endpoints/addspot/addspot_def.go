package addspot

import (
	"net/http"

	"github.com/doniacld/outdoorsight/internal/endpointdef"
	"github.com/doniacld/outdoorsight/internal/spot"
)

// AddSpotMeta holds the endpoint information
var AddSpotMeta = endpointdef.New(
	"addSpot",
	"/spots",
	http.MethodPost,
	http.StatusCreated,
)

// AddSpotRequest is the request structure
var AddSpotRequest spot.Details

// AddSpotResponse holds the response structure
var AddSpotResponse struct{}

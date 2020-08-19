package endpoints

import (
	"net/http"

	"github.com/doniacld/outdoorsight/endpointdef"
)

// GetAPIDocMeta holds the endpoint information
var GetAPIDocMeta = endpointdef.New(
	"getAPIDoc",
	"/doc/api",
	http.MethodGet,
	http.StatusOK,
)

// GetAPIDocRequest is the request structure
var GetAPIDocRequest struct{}

// GetAPIDocResponse holds the response structure
var GetAPIDocResponse []byte

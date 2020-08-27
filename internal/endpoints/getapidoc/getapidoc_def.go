package getapidoc

import (
	"net/http"

	"github.com/doniacld/outdoorsight/internal/endpointdef"
)

const (
	DocAPIFile = "./doc/api/index.html"
)

// GetAPIDocMeta holds the endpoint information
var GetAPIDocMeta = endpointdef.New(
	"getAPIDoc",
	"/doc/api",
	http.MethodGet,
	http.StatusOK,
)

// GetAPIDocRequest is the request structure
type GetAPIDocRequest struct{}

// GetAPIDocResponse holds the response structure
type GetAPIDocResponse []byte

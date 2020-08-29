package errors

import (
	"encoding/json"
	"fmt"
	"github.com/doniacld/outdoorsight/internal/endpointdef"
	"net/http"
)

// OsError represents the internal representation of an error
type OsError struct {
	HTTPCode int    `json:"httpCode"`
	Message  string `json:"message"`
}

func New(errorType int, message string) *OsError {
	return &OsError{HTTPCode: errorType, Message: message}
}

func NewFromError(errorType int, err error, message string) *OsError {
	return &OsError{HTTPCode: errorType, Message: fmt.Sprintf("%s, %s", message, err.Error())}
}

func Wrap(err *OsError, message string) *OsError {
	return &OsError{HTTPCode: err.HTTPCode, Message: fmt.Sprintf("%s: %s", message, err.Message)}
}

// HTTPError encodes the http error into a JSON file
func HTTPError(w http.ResponseWriter, err *OsError) {
	if jsonErr := json.NewEncoder(w).Encode(&err); jsonErr != nil {
		err.Message = jsonErr.Error()
	}

	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeJSON)
	w.WriteHeader(err.HTTPCode)
}

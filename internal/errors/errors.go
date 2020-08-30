package errors

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/doniacld/outdoorsight/internal/endpointdef"
)

// ODSError represents the format of a returned HTTP error
type ODSError struct {
	HTTPCode int    `json:"HTTPCode"`
	Message  string `json:"message"`
}

// Error returns the message of the error
func (err *ODSError) Error() string {
	return err.Message
}

func New(errorType int, message string) *ODSError {
	return &ODSError{HTTPCode: errorType, Message: message}
}

func NewFromError(errorType int, err error, message string) *ODSError {
	return &ODSError{HTTPCode: errorType, Message: fmt.Sprintf("%s, %s", message, err.Error())}
}

func (err *ODSError) Wrap(message string) *ODSError {
	return &ODSError{HTTPCode: err.HTTPCode, Message: fmt.Sprintf("%s: %s", message, err.Message)}
}

// HTTPError encodes the http error into a JSON file
func (err *ODSError) HTTPError(w http.ResponseWriter) {
	if jsonErr := json.NewEncoder(w).Encode(&err); jsonErr != nil {
		err.Message = jsonErr.Error()
	}

	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeJSON)
	w.WriteHeader(err.HTTPCode)
}

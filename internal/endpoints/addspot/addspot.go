package addspot

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"

	"github.com/doniacld/outdoorsight/internal/db"
	"github.com/doniacld/outdoorsight/internal/endpointdef"
)

// AddSpot adds a spot to your list
func AddSpot(w http.ResponseWriter, r *http.Request) {
	// decode the body into spotDetails structure
	var spotDetailsDB db.SpotDetails
	if err := json.NewDecoder(r.Body).Decode(&spotDetailsDB); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// call the database
	osDB := db.New()
	if err := osDB.AddSpot(context.TODO(), spotDetailsDB); err != nil {
		panic(errors.Wrap(err, "unable to add spot"))
	}

	// set the response parameters
	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeJSON)
	w.WriteHeader(AddSpotMeta.SuccessCode())
}

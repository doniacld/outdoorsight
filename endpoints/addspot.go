package endpoints

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/doniacld/outdoorsight/db"
	"github.com/doniacld/outdoorsight/endpointdef"
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

	// TODO DONIA db part should be done elsewhere
	collection := db.ConnectDB()
	_, err := collection.InsertOne(context.TODO(), spotDetailsDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeJSON)
	w.WriteHeader(AddSpotMeta.SuccessCode())
}

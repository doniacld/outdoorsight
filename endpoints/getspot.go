package endpoints

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/doniacld/outdoorsight/db"
	"github.com/doniacld/outdoorsight/endpointdef"
	"github.com/doniacld/outdoorsight/spot"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// GetSpot returns all the details on a given spot
func GetSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	spotName := vars[ParamSpotName]

	// TODO DONIA db part should be done elsewhere

	// connect to the DB
	collection := db.ConnectDB()
	// request the DB with the given spot name
	query := bson.M{
		"name": spotName,
	}
	cursor, err := collection.Find(context.Background(), query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer cursor.Close(context.Background())

	// convert the result which is a cursor in a spotDetails structure
	for cursor.Next(context.Background()) {
		var spotDetails spot.Details
		err := cursor.Decode(&spotDetails)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set the response parameters
		w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeJSON)
		w.WriteHeader(GetSpotMeta.SuccessCode())
		if err := json.NewEncoder(w).Encode(spotDetails); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

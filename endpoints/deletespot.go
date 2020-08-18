package endpoints

import (
	"context"
	"net/http"

	"github.com/doniacld/outdoorsight/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteSpot returns all the details on a given spot
func DeleteSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	spotName := vars[ParamSpotName]

	// TODO DONIA db part should be done elsewhere

	// connect to the DB
	collection := db.ConnectDB()
	// request the DB with the given spot name
	query := bson.M{
		"name": spotName,
	}
	_, err := collection.DeleteOne(context.Background(), query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// set the response parameters
	w.WriteHeader(DeleteSpotMeta.SuccessCode())
}

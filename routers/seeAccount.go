package routers

import (
	"encoding/json"
	"github.com/dest92/Twitty/database"
	"net/http"
)

//SeeProfile shows the profile of a user

func SeeProfile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id") //Extract the id from the body

	if len(ID) < 1 { 
		http.Error(w, "You must send the ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := database.SearchProfile(ID) //Search the user in the database

	if err != nil {
		http.Error(w, "An error occurred while trying to search the user: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json") //Set the header to json
	w.WriteHeader(http.StatusCreated) //Set the status to 201
	json.NewEncoder(w).Encode(profile) //Encode the profile in json and send it to the client
}

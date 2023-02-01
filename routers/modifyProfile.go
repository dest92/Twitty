package routers

import (
	"encoding/json"
	"github.com/dest92/Twitty/database"
	"github.com/dest92/Twitty/models"

	"net/http"
)

// ModifyProfile modify the user profile

func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User //Create user model

	err := json.NewDecoder(r.Body).Decode(&t) //Decode the body of the request

	if err != nil {
		http.Error(w, "Error in the received data: "+err.Error(), 400)
		return
	}

	var status bool
	status, err = database.UpdateProfile(t, UserID)

	if err != nil {
		http.Error(w, "An error occurred while trying to modify the profile: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "An error occurred while trying to modify the profile", 400)
		return
	}

	w.WriteHeader(http.StatusCreated) // 201 status code

}

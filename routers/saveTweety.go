package routers

import (
	"encoding/json"
	"github.com/dest92/Twitty/database"
	"github.com/dest92/Twitty/models"
	"net/http"
	"time"
)

func CreateTweety(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json") // Set the header to json

	var t models.Tweety

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error in the received data: "+err.Error(), 400)
		return
	}

	if len(t.Message) == 0 {
		http.Error(w, "The message is required", 400)
		return
	}

	t.UserID = UserID
	t.Date = time.Now()

	var status bool
	_, status, err = database.CreateTweety(t)

	if err != nil {
		http.Error(w, "An error occurred while trying to register the user: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "An error occurred while trying to register the user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated) // 201

}

package routers

import (
	"github.com/dest92/Twitty/database"
	"net/http"
)

func EraseTweety(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the ID parameter", http.StatusBadRequest)
		return
	}

	err := database.DeleteTweety(ID, UserID)
	if err != nil {
		http.Error(w, "An error occurred while trying to delete the tweet: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json") 
	w.WriteHeader(http.StatusCreated)
}

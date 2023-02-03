package routers

import (
	"github.com/dest92/Twitty/database"
	"github.com/dest92/Twitty/models"
	"net/http"
)

func EraseRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the ID parameter", http.StatusBadRequest)
		return
	}

	var t models.Relation

	t.UserID = UserID
	t.UserRelationID = ID

	status, err := database.DeleteRelation(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to delete the relation: "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "An error occurred while trying to delete the relation: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

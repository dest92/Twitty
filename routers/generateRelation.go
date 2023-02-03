package routers

import (
	"github.com/dest92/Twitty/database"
	"github.com/dest92/Twitty/models"
	"net/http"
)

func GenerateRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "The ID parameter is required", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = UserID
	t.UserRelationID = ID

	if t.UserID == t.UserRelationID {
		http.Error(w, "You cannot follow yourself", 400)
		return
	}

	status, err := database.CreateRelation(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to generate the relation "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "An error occurred while trying to generate the relation ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

package routers

import (
	"encoding/json"
	"github.com/dest92/Twitty/database"
	"github.com/dest92/Twitty/models"
	"net/http"
)

func ConsultRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the ID parameter", http.StatusBadRequest)
		return
	}

	var t models.Relation

	t.UserID = UserID
	t.UserRelationID = ID

	var resp models.RelationResponse

	status, err := database.GetRelation(t)

	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp) // Encode the response
}

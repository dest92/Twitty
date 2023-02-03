package routers

import (
	"encoding/json"
	"github.com/dest92/Twitty/database"
	"net/http"
	"strconv"
)

// ConsultUsers consults the users

func ConsultUsers(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type") // Get the type of user to consult

	if len(typeUser) < 1 {
		http.Error(w, "You must send the type parameter", http.StatusBadRequest)
		return
	}

	page := r.URL.Query().Get("page") // Get the page to consult
	if len(page) < 1 {
		http.Error(w, "You must send the page parameter", http.StatusBadRequest)
		return
	}

	search := r.URL.Query().Get("search") // Get the search to consult

	pagTemp, err := strconv.Atoi(page) // Convert the page to int

	if err != nil {
		http.Error(w, "You must send the page parameter as a greater than zero integer", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp) // Convert the page to int64

	result, status := database.ReadAllUsers(UserID, pag, search, typeUser) // Consult the users

	if !status {
		http.Error(w, "Error reading users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result) // Encode the response

}

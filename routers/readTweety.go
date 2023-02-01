package routers

import (
	"encoding/json"
	"github.com/dest92/Twitty/database"
	"net/http"
	"strconv"
)

func LookTweety(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the parameter ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send the parameter page", http.StatusBadRequest)
		return
	}

	// Convert the page parameter to an int

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "You must send the parameter page", http.StatusBadRequest)
		return
	}

	pag := int64(page)

	result, status := database.ReadTweety(ID, pag) //Results is a slice of pointers to ReadTweety

	if !status {
		http.Error(w, "Error reading the tweetys", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(result) // Encode the result to json and send it to the client

}

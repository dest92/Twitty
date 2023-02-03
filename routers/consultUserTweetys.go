package routers

import (
	"encoding/json"
	"github.com/dest92/Twitty/database"
	"log"
	"net/http"
	"strconv"
)

func ConsultUserTweetys(w http.ResponseWriter, r *http.Request) {

	// Get the page parameter
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		http.Error(w, "Missing 'page' parameter", http.StatusBadRequest)
		return
	}

	// Convert the page to int
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		http.Error(w, "Invalid 'page' parameter, must be a positive integer", http.StatusBadRequest)
		return
	}

	// Consult the user tweetys
	result, status := database.ReadUserTweetys(UserID, int64(page))
	if !status {
		http.Error(w, "Error reading user tweetys", http.StatusBadRequest)
		return
	}

	// Encode the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Error encoding response:", err)
	}
}

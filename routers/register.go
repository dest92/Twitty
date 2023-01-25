package routers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dest92/Twitty/database"
	"github.com/dest92/Twitty/models"
)

// Register is a function to register a new user

func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User

	//! Create a user json model and decode the body into that model
	//! Everything that comes in the body ends up being a json
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error in user data: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	if len(t.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters", 400)
		return
	}

	//Check if the password has a uppercase
	var hasUpper bool
	for _, letter := range t.Password {
		if strings.ToUpper(string(letter)) == string(letter) {
			hasUpper = true
			break
		}
	}

	if !hasUpper {
		http.Error(w, "Password must have at least one uppercase letter", 400)
		return
	}
	

	_, found, _ := database.UserExists(t.Email)

	if found {
		http.Error(w, "There is already a registered user with that email", 400)
		return
	}

	_, status, err := database.CreateUser(t)

	if err != nil {
		http.Error(w, "Error registering user: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Error registering user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

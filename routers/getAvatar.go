package routers

import (
	"github.com/dest92/Twitty/database"
	"io"
	"net/http"
	"os"
)

func GetAvatar(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "You must send the ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := database.SearchProfile(ID) //Search the user in the database

	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + profile.Avatar) //Open the file

	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile) //Copy the file to the response writer

	if err != nil {
		http.Error(w, "Error while copying the image", http.StatusBadRequest)
		return
	}

}

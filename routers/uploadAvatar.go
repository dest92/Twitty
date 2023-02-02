package routers

import (
	"github.com/dest92/Twitty/database"
	"github.com/dest92/Twitty/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar") //Get the file from the request and save it in the variable file and the name in the variable handler

	if err != nil {
		http.Error(w, "Error while uploading the image: "+err.Error(), http.StatusBadRequest)
		return
	}

	var extension = filepath.Ext(handler.Filename) //Get the extension of the file

	var archive string = "uploads/avatars/" + UserID + extension

	f, err := os.OpenFile(archive, os.O_WRONLY|os.O_CREATE, 0666) //Create the file in the server

	if err != nil {
		http.Error(w, "Error while uploading the image: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file) //Copy the file to the server

	if err != nil {
		http.Error(w, "Error while copying the image: "+err.Error(), http.StatusBadRequest)
		return
	}

	//Save the path of the image in the database

	var user models.User
	var status bool

	user.Avatar = UserID + extension
	status, err = database.UpdateProfile(user, UserID)

	if err != nil || !status {
		http.Error(w, "Error while saving the image in the database: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

package routers

import (
	"encoding/json"
	"github.com/dest92/Twitty/database"
	"github.com/dest92/Twitty/jwt"
	"github.com/dest92/Twitty/models"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//Set header
	w.Header().Add("content-type", "application/json") //Set the header to json
	//ResponseWriter return the response to the client

	var t models.User

	//Decode the request body
	err := json.NewDecoder(r.Body).Decode(&t) //Decode the request body and put it in the t variable

	//Login only accept email and password

	if err != nil {
		http.Error(w, "User or password incorrect: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	//Compare the email and password with the database
	document, exists := database.TryLogin(t.Email, t.Password)

	if exists == false {
		http.Error(w, "User or password incorrect", 400)
		return
	}

	//Create a JWT

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "Error generating the token: "+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	} //Return to navigator the token

	w.Header().Set("Content-Type", "application/json") //Set the header to json
	w.WriteHeader(http.StatusCreated)                  //Set the status code (200 or 201)
	json.NewEncoder(w).Encode(resp)                    //Encode the response and send it to the client

	//Save cookie

	expirationTime := time.Now().Add(24 * time.Hour) //Set the expiration time to 24 hours
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}

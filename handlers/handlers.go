package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors" //Permissions
	"log"
	"github.com/dest92/Twitty/middlew"
	"github.com/dest92/Twitty/routers"
	"net/http"
	"os"
)

//Set port, handler and listen to server

func Handlers() {
	//! Mux captures the http
	router := mux.NewRouter()

	//! Routes
	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/seeprofile", middlew.CheckJWT(routers.SeeProfile)).Methods("POST")


	


	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	//! Handler object
	handler := cors.AllowAll().Handler(router) //Permissions to everything

	//! Listen to server
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //Listen to port
}

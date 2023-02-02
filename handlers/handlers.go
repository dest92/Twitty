package handlers

import (
	"github.com/dest92/Twitty/middlew"
	"github.com/dest92/Twitty/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors" //Permissions
	"log"
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
	router.HandleFunc("/seeprofile", middlew.CheckJWT(routers.SeeProfile)).Methods("GET")
	router.HandleFunc("/updateProfile", middlew.CheckJWT(routers.ModifyProfile)).Methods("PUT")
	router.HandleFunc("/tweety", middlew.CheckJWT(routers.CreateTweety)).Methods("POST")
	router.HandleFunc("/readTweety", middlew.CheckJWT(routers.LookTweety)).Methods("GET")
	router.HandleFunc("/eraseTweety", middlew.CheckJWT(routers.EraseTweety)).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	//! Handler object
	handler := cors.AllowAll().Handler(router) //Permissions to everything

	//! Listen to server
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //Listen to port
}

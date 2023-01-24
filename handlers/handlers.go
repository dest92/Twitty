package handlers

import (
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

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	//! Handler object
	handler := cors.AllowAll().Handler(router)  //Permissions to everything

	//! Listen to server
	 log.Fatal(http.ListenAndServe(":"+PORT,handler)) //Listen to port
}

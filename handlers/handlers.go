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
	router.HandleFunc("/seeprofile", middlew.CheckDB(middlew.CheckJWT(routers.SeeProfile))).Methods("GET")
	router.HandleFunc("/updateProfile", middlew.CheckDB(middlew.CheckJWT(routers.ModifyProfile))).Methods("PUT")

	router.HandleFunc("/tweety", middlew.CheckDB(middlew.CheckJWT(routers.CreateTweety))).Methods("POST")
	router.HandleFunc("/readTweety", middlew.CheckDB(middlew.CheckJWT(routers.LookTweety))).Methods("GET")
	router.HandleFunc("/eraseTweety", middlew.CheckDB(middlew.CheckJWT(routers.EraseTweety))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlew.CheckDB(middlew.CheckJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/uploadBanner", middlew.CheckDB(middlew.CheckJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/getBanner", middlew.CheckDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/generateRelation", middlew.CheckDB(middlew.CheckJWT(routers.GenerateRelation))).Methods("POST")
	router.HandleFunc("/eraseRelation", middlew.CheckDB(middlew.CheckJWT(routers.EraseRelation))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	//! Handler object
	handler := cors.AllowAll().Handler(router) //Permissions to everything

	//! Listen to server
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //Listen to port
}

package main

import (
	db "github.com/dest92/Twitty/database"
	hd "github.com/dest92/Twitty/handlers"
	"log"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Error connecting to database")
		return
	}
	log.Println("Starting the application...")
	hd.Handlers()
}

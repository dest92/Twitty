package main

import(
	"log"
	db "github.com/dest92/Twitty/database"
	hd "github.com/dest92/Twitty/handlers"
)

func main() {
	log.Println("Starting the application...")
	if(db.CheckConnection() == 0){
		log.Fatal("Error connecting to database")
		return
	}
	hd.Handlers()
}


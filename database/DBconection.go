package database

import (
	"context"
	//"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// !Execute the connection to database
// !Return to Mongo the connection
var MongoCN = conectDB()

// !Connect to URl

// !Load the .env file
var load = godotenv.Load()

var clientOptions = options.Client().ApplyURI(os.Getenv("HOST_URL"))

// !Connect to database
func conectDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	//Asign new valor to err
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connected to database")

	return client

}

// !Check the connection
func CheckConnection() int {

	if load != nil {
		log.Fatal(load)
		return 0
	}

	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1

}

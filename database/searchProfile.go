package database

import (
	"context"
	"github.com/dest92/Twitty/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

//SearchProfile searches for a user in the database by ID

func SearchProfile(ID string) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("Twitty")
	col := db.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID) //Convert the ID to an objectID

	condition := bson.M{"_id": objID} // Compare the ID with the ID in the database

	err := col.FindOne(ctx, condition).Decode(&profile) //Decode the result, convert in json and put it in the profile variable

	if err != nil {
		log.Fatal("Registry was not found" + err.Error())
		return profile, err
	}

	profile.Password = ""

	return profile, nil

}

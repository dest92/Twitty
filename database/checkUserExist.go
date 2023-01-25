package database

import (
	"context"
	"github.com/dest92/Twitty/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UserExists(email string) (models.User, bool, string) {

	//Check if the user already exists
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitty")

	col := db.Collection("users")

	condition := bson.M{"email": email} //M = Map

	var result models.User //User model

	err := col.FindOne(ctx, condition).Decode(&result) //Decode the result, convert in json and put it in the result variable

	ID := result.ID.Hex() //Get the ID and convert to string

	if err != nil {
		return result, false, ID
	}

	return result, true, ID //If the user exists, return true

}

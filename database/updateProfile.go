package database

import (
	"context"
	"github.com/dest92/Twitty/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Update the user profile
func UpdateProfile(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitty")
	col := db.Collection("users")

	profile := make(map[string]interface{}) //Create a map to store the values to update

	//Check

	//If the name is not empty, add it to the map

	if len(u.Name) > 0 {
		profile["name"] = u.Name
	}

	//If the last name is not empty, add it to the map

	if len(u.LastName) > 0 {
		profile["lastName"] = u.LastName
	}

	//If the birth date is not empty, add it to the map

	if !u.BirthDate.IsZero() || u.BirthDate != (time.Time{}) {
		profile["birthDate"] = u.BirthDate
	}

	//If the biography is not empty, add it to the map

	if len(u.Biography) > 0 {
		profile["biography"] = u.Biography
	}

	//If the location is not empty, add it to the map

	if len(u.Location) > 0 {
		profile["location"] = u.Location
	}

	//If the website is not empty, add it to the map

	if len(u.WebSite) > 0 {
		profile["website"] = u.WebSite
	}

	//Set fields to update

	updtString := bson.M{
		"$set": profile,
	}

	//Convert the ID to an ObjectID

	objID, _ := primitive.ObjectIDFromHex(ID)

	//Create a filter to find the user by ID

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)

	if err != nil {
		return false, err
	}

	return true, nil

}

package database

import (
	"context"
	"github.com/dest92/Twitty/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// ReadAllUsers reads all the users from the database

func ReadAllUsers(ID string, page int64, search string, typeUser string) ([]*models.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitty")
	col := db.Collection("users")

	var results []*models.User //To send to http

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20) // Skip the first 20 results
	findOptions.SetLimit(20)             // Limit the results to 20

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search}, // (?i) is to ignore case
	}

	cur, err := col.Find(ctx, query, findOptions) // Find the users that match the query

	if err != nil {
		log.Println(err.Error())
		return results, false
	}

	var found, include bool // To check if the user is found and if it should be included

	for cur.Next(ctx) {
		var u models.User
		err := cur.Decode(&u)

		if err != nil {
			log.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationID = u.ID.Hex()

		include = false // To check if the user should be included

		found, _ = GetRelation(r) // Check if the user is already following the user

		//New is the users that the user is not following

		if typeUser == "new" && !found { // If the user is not found and the type is new, include the user
			include = true
		}

		//Following is the users that the user is following

		if typeUser == "following" && found { // If the user is found and the type is following, include the user
			include = true
		}

		if r.UserRelationID == ID { // If the user is the same as the user that is logged in, don't include the user
			include = false
		}

		if include {
			//Erase the fields that we don't want to send to the frontend
			u.Password = ""
			u.Biography = ""
			u.WebSite = ""
			u.Location = ""
			u.Banner = ""
			u.Email = ""
			results = append(results, &u)
		}
	}

	err = cur.Err()

	if err != nil {
		log.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)
	return results, true
}

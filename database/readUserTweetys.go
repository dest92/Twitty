package database

import (
	"context"
	"github.com/dest92/Twitty/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// ReadUserTweetys reads all the tweetys from the database

func ReadUserTweetys(ID string, page int64) ([]*models.TweetysResponse, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitty")
	col := db.Collection("relation")

	skip := (page - 1) * 20 // Skip the first 20 results

	conditions := make([]bson.M, 0)                                         // To store the conditions
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}}) // Match the user id
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{ // Join the relation collection with the tweety collection
			"from":         "tweety",
			"localField":   "userRelationId",
			"foreignField": "userid",
			"as":           "tweety",
		}})

	conditions = append(conditions, bson.M{"$unwind": "$tweety"})               // Unwind the tweety array
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweety.date": -1}}) // Sort the tweety by date
	conditions = append(conditions, bson.M{"$skip": skip})                      // Skip the first 20 results
	conditions = append(conditions, bson.M{"$limit": 20})                       // Limit the results to 20

	cursor, err := col.Aggregate(ctx, conditions) // Aggregate the conditions

	if err != nil {
		return nil, false
	}

	var result []*models.TweetysResponse // To send to http

	err = cursor.All(ctx, &result) // Get all the results

	if err != nil {
		return result, false
	}

	return result, true

}

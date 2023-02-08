package database

import (
	"context"
	"github.com/dest92/Twitty/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

// ReadUserTweetys reads all the tweetys from the database

func ReadUserTweetys(IDUser string, page int) ([]models.TweetysResponse, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitty")
	col := db.Collection("relation")

	skip := (page - 1) * 20

	conditions := []bson.M{
		{"$match": bson.M{"userId": IDUser}},
		{"$lookup": bson.M{"from": "tweetys", "localField": "userRelationId", "foreignField": "userid", "as": "tweetys"}},
		{"$unwind": "$tweetys"},
		{"$sort": bson.M{"tweetys.date": -1}},
		{"$skip": skip},
		{"$limit": 20},
	}

	cursor, err := col.Aggregate(ctx, conditions)
	if err != nil {
		log.Println("Error adding conditions: ", err)
		return nil, false
	}

	var result []models.TweetysResponse
	err = cursor.All(ctx, &result)
	if err != nil {
		log.Println("Error getting results: ", err)
		return nil, false
	}

	if len(result) == 0 {
		log.Println("No results found")
		return nil, false
	}

	return result, true
}

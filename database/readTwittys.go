package database

import (
	"context"
	"github.com/dest92/Twitty/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ReadTweety(ID string, page int64) ([]*models.ReadTweety, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitty")
	col := db.Collection("tweety")

	var results []*models.ReadTweety

	condition := bson.M{
		"userid": ID,
	}

	// Set the options to the query
	options := options.Find()
	options.SetLimit(20)                              // Limit the results to 20
	options.SetSort(bson.D{{Key: "date", Value: -1}}) // Sort the results by date in descending order
	options.SetSkip((page - 1) * 20)                  // Skip the first 20 results

	cursor, err := col.Find(ctx, condition, options) //Cursor is a pointer to the results
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	//Foreach document in the cursor decode it and append it to the results

	for cursor.Next(context.TODO()) { //Iterate over the cursor and decode each document
		var register models.ReadTweety
		err := cursor.Decode(&register) //Decode the document
		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}

	return results, true

}

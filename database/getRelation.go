package database

import (
	"context"
	"github.com/dest92/Twitty/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

// GetRelation checks if there is a relation between two users

func GetRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitty")
	col := db.Collection("relation")

	condition := bson.M{
		"userid":         t.UserID,
		"userrelationid": t.UserRelationID,
	}

	var result models.Relation
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}

package database

import (
	"context"
	"time"

	"github.com/dest92/Twitty/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //Cancel the context with the timeout

	db := MongoCN.Database("twitty")

	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil

}

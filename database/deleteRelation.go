package database

import (
	"context"
	"github.com/dest92/Twitty/models"
	"time"
)

// DeleteRelation deletes a relation from the database

func DeleteRelation(r models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitty")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, r)
	if err != nil {
		return false, err
	}
	return true, nil
}

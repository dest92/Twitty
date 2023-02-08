package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tweety struct {
	UserID  string    `bson:"userid,omitempty" json:"userid,omitempty"`
	Message string    `bson:"message,omitempty" json:"message,omitempty"`
	Date    time.Time `bson:"date,omitempty" json:"date,omitempty"`
}

type ReadTweety struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userid,omitempty" json:"userid,omitempty"`
	Message string             `bson:"message,omitempty" json:"message,omitempty"`
	Date    time.Time          `bson:"date,omitempty" json:"date,omitempty"`
}

type TweetysResponse struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid,omitempty" json:"userid,omitempty"`
	UserRelationID string             `bson:"userRelationId,omitempty" json:"userRelationId,omitempty"`
	Tweetyrp       struct {
		ID      string    `bson:"_id" json:"_id,omitempty"`
		Message string    `bson:"message,omitempty" json:"message,omitempty"`
		Date    time.Time `bson:"date,omitempty" json:"date,omitempty"`
	} `bson:"tweetys,omitempty" json:"tweetys,omitempty"`
}

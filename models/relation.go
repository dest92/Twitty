package models

type Relation struct {
	UserID         string `bson:"userId" json:"userId,omitempty"`
	UserRelationID string `bson:"userRelationId" json:"userRelationId,omitempty"`
}

type RelationResponse struct {
	Status bool `json:"status"`
}

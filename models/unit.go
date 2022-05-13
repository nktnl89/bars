package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Unit struct {
	CreatorID   int                `json:"creator_id" bson:"creator_id"`
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type UnitUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        Unit
}

type UnitDelete struct {
	DeletedCount int64 `json:"deletedCount"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {
	ID      primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string               `json:"name" bson:"name"`
	UnitIds []primitive.ObjectID `json:"unit_ids" bson:"unit_ids"`
}

type GroupUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        Group
}

type GroupDelete struct {
	DeletedCount int64 `json:"deletedCount"`
}

type GroupFetch struct {
	Group Group  `json:"group"`
	Units []Unit `json:"units"`
}

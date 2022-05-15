package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name" json:"name" validate:"required,min=2,max=100"`
	Password      string             `bson:"password" json:"password" validate:"required,min=3""`
	Token         string             `bson:"token" json:"token"`
	Refresh_token string             `bson:"refresh_token" json:"refresh_token"`
	Created_at    time.Time          `bson:"created_at" json:"created_at"`
	Updated_at    time.Time          `bson:"updated_at" json:"updated_at"`
}

package database

import (
	"bars/bars/models"
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserInterface interface {
	Insert(models.User) (models.User, error)
	Get(string) (models.User, error)
	Update(string, interface{}) (models.User, error)
}

type UserClient struct {
	Ctx context.Context
	Col *mongo.Collection
}

func (c *UserClient) Insert(docs models.User) (models.User, error) {
	user := models.User{}
	res, err := c.Col.InsertOne(c.Ctx, docs)
	if err != nil {
		return user, err
	}
	res.InsertedID.(primitive.ObjectID).Hex()
	return c.Get(docs.Name)
}

func (c *UserClient) Get(name string) (models.User, error) {
	user := models.User{}

	err := c.Col.FindOne(c.Ctx, bson.M{"name": name}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (c *UserClient) Update(name string, update interface{}) (models.User, error) {
	result := models.User{}

	user, err := c.Get(name)
	if err != nil {
		return result, err
	}
	var exist map[string]interface{}
	b, err := json.Marshal(user)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &exist)
	change := update.(map[string]interface{})
	for k := range change {
		if change[k] == exist[k] {
			delete(change, k)
		}
	}
	if len(change) == 0 {
		return result, nil
	}
	res, err := c.Col.UpdateOne(c.Ctx, bson.M{"name": name}, bson.M{"$set": change})
	if err != nil {
		return result, err
	}
	//TODO надо убрать както res
	res.UpsertedID.(primitive.ObjectID).Hex()

	newUser, err := c.Get(name)
	if err != nil {
		return result, err
	}

	return newUser, nil
}

package database

import (
	"bars/bars/models"
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GroupInterface interface {
	Insert(models.Group) (models.Group, error)
	Update(string, interface{}) (models.GroupUpdate, error)
	Delete(string) (models.GroupDelete, error)
	Get(string) (models.Group, error)
	FindAll() ([]models.Group, error)
	FetchGroup(string) (models.GroupFetch, error)
	FetchAllGroups() ([]models.GroupFetch, error)
}

type GroupClient struct {
	Ctx context.Context
	Col *mongo.Collection
}

func (c *GroupClient) Insert(docs models.Group) (models.Group, error) {
	group := models.Group{}

	res, err := c.Col.InsertOne(c.Ctx, docs)
	if err != nil {
		return group, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return c.Get(id)
}

func (c *GroupClient) Update(id string, update interface{}) (models.GroupUpdate, error) {
	result := models.GroupUpdate{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, nil
	}

	group, err := c.Get(id)
	if err != nil {
		return result, err
	}
	var exist map[string]interface{}
	b, err := json.Marshal(group)
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
	res, err := c.Col.UpdateOne(c.Ctx, bson.M{"_id": _id}, bson.M{"$set": change})
	if err != nil {
		return result, nil
	}
	newGroup, err := c.Get(id)
	if err != nil {
		return result, err
	}
	result.ModifiedCount = res.ModifiedCount
	result.Result = newGroup
	return result, nil
}

func (c *GroupClient) Get(id string) (models.Group, error) {
	group := models.Group{}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return group, err
	}
	err = c.Col.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&group)
	if err != nil {
		return group, err
	}
	return group, nil

}

func (c *GroupClient) Delete(id string) (models.GroupDelete, error) {
	result := models.GroupDelete{
		DeletedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, nil
	}

	res, err := c.Col.DeleteOne(c.Ctx, bson.M{"_id": _id})
	if err != nil {
		return result, err
	}
	result.DeletedCount = res.DeletedCount
	return result, nil
}

func (c *GroupClient) FindAll() ([]models.Group, error) {
	groups := []models.Group{}

	cursor, err := c.Col.Find(c.Ctx, bson.M{})
	if err != nil {
		return groups, err
	}

	for cursor.Next(c.Ctx) {
		row := models.Group{}
		cursor.Decode(&row)
		groups = append(groups, row)
	}

	return groups, nil
}

func (c *GroupClient) FetchGroup(id string) (models.GroupFetch, error) {
	res := models.GroupFetch{}

	group := models.Group{}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return res, err
	}
	err = c.Col.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&group)
	if err != nil {
		return res, err
	}
	res.Group = group

	units, err := findAllUnitsByIds(c, group.UnitIds)
	if err != nil {
		return res, err
	}
	res.Units = units

	return res, nil
}

func (c *GroupClient) FetchAllGroups() ([]models.GroupFetch, error) {
	res := []models.GroupFetch{}
	groups := []models.Group{}

	cursor, err := c.Col.Find(c.Ctx, bson.M{})
	if err != nil {
		log.Println(cursor)
		return res, err
	}
	for cursor.Next(c.Ctx) {
		row := models.Group{}
		cursor.Decode(&row)
		groups = append(groups, row)
	}

	for _, g := range groups {
		var units, err = findAllUnitsByIds(c, g.UnitIds)
		if err != nil {
			return res, err
		}
		res = append(res, models.GroupFetch{
			Group: g,
			Units: units,
		})
	}

	return res, nil
}

func findAllUnitsByIds(c *GroupClient, ids []primitive.ObjectID) ([]models.Unit, error) {
	var units = []models.Unit{}
	if len(ids) == 0 {
		return units, nil
	}
	cursor, err := c.Col.Find(c.Ctx, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return units, err
	}

	for cursor.Next(c.Ctx) {
		row := models.Unit{}
		cursor.Decode(&row)
		units = append(units, row)
	}
	return units, nil
}

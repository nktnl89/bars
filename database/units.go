package database

import (
	"bars/bars/models"
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UnitInterface interface {
	Insert(models.Unit) (models.Unit, error)
	Update(string, interface{}) (models.UnitUpdate, error)
	Delete(string) (models.UnitDelete, error)
	Get(string) (models.Unit, error)
	FindAll() ([]models.Unit, error)
	Search(interface{}) ([]models.Unit, error)
}

type UnitClient struct {
	Ctx context.Context
	Col *mongo.Collection
}

func (c *UnitClient) Insert(docs models.Unit) (models.Unit, error) {
	unit := models.Unit{}

	res, err := c.Col.InsertOne(c.Ctx, docs)
	if err != nil {
		return unit, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return c.Get(id)
}

func (c *UnitClient) Update(id string, update interface{}) (models.UnitUpdate, error) {
	result := models.UnitUpdate{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, nil
	}

	unit, err := c.Get(id)
	if err != nil {
		return result, err
	}
	var exist map[string]interface{}
	b, err := json.Marshal(unit)
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
	newUnit, err := c.Get(id)
	if err != nil {
		return result, err
	}
	result.ModifiedCount = res.ModifiedCount
	result.Result = newUnit
	return result, nil
}

func (c *UnitClient) Get(id string) (models.Unit, error) {
	unit := models.Unit{}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return unit, err
	}
	err = c.Col.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&unit)
	if err != nil {
		return unit, err
	}
	return unit, nil

}

func (c *UnitClient) Delete(id string) (models.UnitDelete, error) {
	//TODO: не забудь сделать удаление из группы идшников удаляемых юнитов
	result := models.UnitDelete{
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

func (c *UnitClient) FindAll() ([]models.Unit, error) {
	units := []models.Unit{}

	cursor, err := c.Col.Find(c.Ctx, bson.M{})
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

func (c *UnitClient) Search(filter interface{}) ([]models.Unit, error) {
	units := []models.Unit{}
	if filter == nil {
		filter = bson.M{}
	}

	cursor, err := c.Col.Find(c.Ctx, filter)
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

package database

import (
	"bars/bars/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FetchedInterface interface {
	FetchGroup(string) (models.GroupFetch, error)
	FetchAllGroups() ([]models.GroupFetch, error)
}

type FetchedCollectionClient struct {
	Ctx      context.Context
	UnitCol  *mongo.Collection
	GroupCol *mongo.Collection
}

func (c *FetchedCollectionClient) FetchGroup(id string) (models.GroupFetch, error) {
	res := models.GroupFetch{}

	group := models.Group{}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return res, err
	}
	err = c.GroupCol.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&group)
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

func (c *FetchedCollectionClient) FetchAllGroups() ([]models.GroupFetch, error) {
	res := []models.GroupFetch{}
	groups := []models.Group{}

	cursor, err := c.GroupCol.Find(c.Ctx, bson.M{})
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

func findAllUnitsByIds(c *FetchedCollectionClient, ids []primitive.ObjectID) ([]models.Unit, error) {
	var units = []models.Unit{}
	if len(ids) == 0 {
		return units, nil
	}
	cursor, err := c.UnitCol.Find(c.Ctx, bson.M{"_id": bson.M{"$in": ids}})
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

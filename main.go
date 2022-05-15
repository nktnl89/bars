package main

import (
	"bars/bars/config"
	"bars/bars/controllers"
	"bars/bars/database"
	"bars/bars/helpers"
	"context"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()
	db := database.ConnectDB(ctx, conf.Mongo)

	unitClient := &database.UnitClient{
		Col: db.Collection("Unit"),
		Ctx: ctx,
	}
	groupClient := &database.GroupClient{
		Col: db.Collection("Group"),
		Ctx: ctx,
	}
	userClient := &database.UserClient{
		Col: db.Collection("User"),
		Ctx: ctx,
	}
	fetchedClient := &database.FetchedCollectionClient{
		GroupCol: db.Collection("Group"),
		UnitCol:  db.Collection("Unit"),
		Ctx:      ctx,
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.POST("/users/signup", controllers.SignUp(userClient))
	router.POST("/users/login", controllers.Login(userClient))

	router.Use(helpers.Authentication())

	router.POST("/units", controllers.InsertUnit(unitClient))
	router.DELETE("/units/{id}", controllers.DeleteUnit(unitClient))
	router.PATCH("/units/{id}", controllers.UpdateUnit(unitClient))
	router.GET("/units/{id}", controllers.GetUnit(unitClient))
	router.GET("/units", controllers.FindAllUnit(unitClient))

	router.POST("/groups", controllers.InsertGroup(groupClient))
	router.DELETE("/groups/{id}", controllers.DeleteGroup(groupClient))
	router.PATCH("/groups/{id}", controllers.UpdateGroup(groupClient))
	router.GET("/groups/{id}", controllers.GetGroup(groupClient))
	router.GET("/groups", controllers.FindAllGroup(groupClient))

	router.GET("/groups/fetch/all", controllers.FetchAllGroup(fetchedClient))
	router.GET("/groups/fetch/{id}", controllers.FetchGroup(fetchedClient))

	router.Run(":8080")
}

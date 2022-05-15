package main

import (
	"bars/bars/config"
	"bars/bars/controllers"
	"bars/bars/database"
	"context"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const userkey = "user"

var secret = []byte("secret")

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}

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
	r := gin.New()

	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore(secret)))

	r.POST("/login", controllers.Login(userClient))
	r.POST("/signup", controllers.SignUp(userClient))
	r.GET("/logout", controllers.Logout())

	private := r.Group("/api")
	private.Use(AuthRequired)
	{
		private.POST("/units", controllers.InsertUnit(unitClient))
		private.DELETE("/units/{id}", controllers.DeleteUnit(unitClient))
		private.PATCH("/units/{id}", controllers.UpdateUnit(unitClient))
		private.GET("/units/{id}", controllers.GetUnit(unitClient))
		private.GET("/units", controllers.FindAllUnit(unitClient))

		private.POST("/groups", controllers.InsertGroup(groupClient))
		private.DELETE("/groups/{id}", controllers.DeleteGroup(groupClient))
		private.PATCH("/groups/{id}", controllers.UpdateGroup(groupClient))
		private.GET("/groups/{id}", controllers.GetGroup(groupClient))
		private.GET("/groups", controllers.FindAllGroup(groupClient))

		private.GET("/groups/fetch/all", controllers.FetchAllGroup(fetchedClient))
		private.GET("/groups/fetch/{id}", controllers.FetchGroup(fetchedClient))
	}
	r.Run(":8080")
}

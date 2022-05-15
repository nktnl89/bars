package controllers

import (
	"bars/bars/database"
	"bars/bars/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertGroup(db database.GroupInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		group := models.Group{}

		if err := c.BindJSON(&group); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(group)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		res, err := db.Insert(group)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func GetGroup(db database.GroupInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Request.URL.Query().Get("id")

		res, err := db.Get(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func FindAllGroup(db database.GroupInterface) gin.HandlerFunc {
	return func(c *gin.Context) {

		res, err := db.FindAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func DeleteGroup(db database.GroupInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Request.URL.Query().Get("id")

		res, err := db.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func UpdateGroup(db database.GroupInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Request.URL.Query().Get("id")

		group := models.Group{}
		if err := c.BindJSON(&group); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(group)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		res, err := db.Update(id, group)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

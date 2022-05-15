package controllers

import (
	"bars/bars/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchAllGroup(db database.FetchedInterface) gin.HandlerFunc {

	return func(c *gin.Context) {

		res, err := db.FetchAllGroups()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func FetchGroup(db database.FetchedInterface) gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Request.URL.Query().Get("id")

		res, err := db.FetchGroup(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, res)
	}
}

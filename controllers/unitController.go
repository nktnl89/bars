package controllers

import (
	"bars/bars/database"
	"bars/bars/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUnit(db database.UnitInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Request.URL.Query().Get("id")

		res, err := db.Delete(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unit item was not deleted"})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func FindAllUnit(db database.UnitInterface) gin.HandlerFunc {
	return func(c *gin.Context) {

		res, err := db.FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can not find all unit items"})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func GetUnit(db database.UnitInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Request.URL.Query().Get("id")

		res, err := db.Get(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can not find unit item"})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

//todo закончить тут
func InsertUnit(db database.UnitInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		unit := models.Unit{}

		if err := c.BindJSON(&unit); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(unit)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		res, err := db.Insert(unit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

func SearchUnits(db database.UnitInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter interface{}
		query := c.Request.URL.Query().Get("q")

		if query != "" {
			err := json.Unmarshal([]byte(query), &filter)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		res, err := db.Search(filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func UpdateUnit(db database.UnitInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Request.URL.Query().Get("id")

		unit := models.Unit{}

		if err := c.BindJSON(&unit); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(unit)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		res, err := db.Update(id, unit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

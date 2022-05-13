package handlers

import (
	"bars/bars/database"
	"bars/bars/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func InsertGroup(db database.GroupInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group := models.Group{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = json.Unmarshal(body, &group)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := db.Insert(group)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, res)
	}
}

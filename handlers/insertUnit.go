package handlers

import (
	"bars/bars/database"
	"bars/bars/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func InsertUnit(db database.UnitInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		unit := models.Unit{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = json.Unmarshal(body, &unit)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := db.Insert(unit)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, res)
	}
}

func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

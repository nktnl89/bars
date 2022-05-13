package handlers

import (
	"bars/bars/database"
	"encoding/json"
	"net/http"
)

func SearchUnits(db database.UnitInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filter interface{}
		query := r.URL.Query().Get("q")

		if query != "" {
			err := json.Unmarshal([]byte(query), &filter)
			if err != nil {
				WriteResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		res, err := db.Search(filter)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, res)
	}
}

package handlers

import (
	"bars/bars/database"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteUnit(db database.UnitInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		res, err := db.Delete(id)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, res)
	}
}

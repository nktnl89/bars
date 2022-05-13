package handlers

import (
	"bars/bars/database"
	"net/http"
)

func FindAllGroup(db database.GroupInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res, err := db.FindAll()
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, res)
	}
}

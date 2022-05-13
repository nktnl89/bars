package main

import (
	"bars/bars/config"
	"bars/bars/database"
	"bars/bars/handlers"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()
	db := database.ConnectDB(ctx, conf.Mongo)

	collection := db.Collection(conf.Mongo.Collection)
	unitClient := &database.UnitClient{
		Col: collection,
		Ctx: ctx,
	}
	groupClient := &database.GroupClient{
		Col: collection,
		Ctx: ctx,
	}

	r := mux.NewRouter()

	r.HandleFunc("/units", handlers.InsertUnit(unitClient)).Methods("POST")
	r.HandleFunc("/units/{id}", handlers.DeleteUnit(unitClient)).Methods("DELETE")
	r.HandleFunc("/units/{id}", handlers.UpdateUnit(unitClient)).Methods("PATCH")
	r.HandleFunc("/units/{id}", handlers.GetUnit(unitClient)).Methods("GET")
	r.HandleFunc("/units", handlers.FindAllUnit(unitClient)).Methods("GET")

	r.HandleFunc("/groups", handlers.InsertGroup(groupClient)).Methods("POST")
	r.HandleFunc("/groups/{id}", handlers.DeleteGroup(groupClient)).Methods("DELETE")
	r.HandleFunc("/groups/{id}", handlers.UpdateGroup(groupClient)).Methods("PATCH")
	r.HandleFunc("/groups/{id}", handlers.GetGroup(groupClient)).Methods("GET")
	r.HandleFunc("/groups", handlers.FindAllGroup(groupClient)).Methods("GET")

	r.HandleFunc("/groups/fetch/all", handlers.FetchAllGroup(groupClient)).Methods("GET")
	r.HandleFunc("/groups/fetch/{id}", handlers.FetchGroup(groupClient)).Methods("GET")

	http.ListenAndServe(":8080", r)
}

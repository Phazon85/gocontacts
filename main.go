package main

import (
	"github.com/gorilla/mux"
	"github.com/phazon85/go_contacts/handler"
	"github.com/phazon85/go_contacts/services"
	"github.com/phazon85/multisql"
)

const (
	configFile = "dev.yaml"
	driverName = "postgres"
)

func main() {

	//load DB connection
	db := multisql.NewDBObject(configFile, driverName)

	//load database services
	init := services.InitDB(db)

	//intializes HTTP router
	handler := handler.NewEntryHandler(init)

	r := mux.NewRouter()
	r.HandleFunc("/entry", handler.HandleGetEntries).Methods("GET")
}

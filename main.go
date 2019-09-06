package main

import (
	"log"
	"net/http"

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
	r.HandleFunc("/entry/{id:[0-9]+}", handler.HandleGetEntriesByID).Methods("GET")
	r.HandleFunc("/entry", handler.HandleAddEntry).Methods("POST")
	r.HandleFunc("/entry", handler.HandleUpdateEntry).Methods("PUT")
	r.HandleFunc("/entry", handler.HandleDeleteEntry).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phazon85/go_contacts/services"
)

func encodeJSON(w http.ResponseWriter, v interface{}) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Panicf("Error encoding JSON")
	}
}

//EntryHandler implements the Actions interface
type EntryHandler struct {
	Service services.Actions
}

//NewEntryHandler returns a new EntryHandler
func NewEntryHandler(svc services.Actions) *EntryHandler {
	return &EntryHandler{
		Service: svc,
	}
}

//HandleGetEntries handles getting all entries in contacts DB
func (e *EntryHandler) HandleGetEntries(w http.ResponseWriter, r *http.Request) {
	res, err := e.Service.AllEntries()
	if err != nil {
		log.Printf("Error getting all entries: %s", err.Error())
	}
	encodeJSON(w, res)
}

//HandleGetEntriesByID handles getting a specific entry by id
func (e *EntryHandler) HandleGetEntriesByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res, err := e.Service.EntryByID(vars["id"])
	if err != nil {
		log.Printf("Error getting todo by id: %s", err.Error())
	}
	w.
		encodeJSON(w, res)
}

package handler

import (
	"encoding/json"
	"log"
	"net/http"

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

//HandleGetEntries get all entries in contacts DB
func (e *EntryHandler) HandleGetEntries(w http.ResponseWriter, r *http.Request) {
	res, err := e.Service.AllEntries()
	if err != nil {
		log.Printf("Error getting all entries: %s", err.Error())
	}
	encodeJSON(w, res)
}

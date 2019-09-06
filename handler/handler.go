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

func decodeAndValidate(r *http.Request, v services.Validation) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()

	return v.Validate()
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
	encodeJSON(w, res)
}

//HandleAddEntry handles adding an entry to the DB
func (e *EntryHandler) HandleAddEntry(w http.ResponseWriter, r *http.Request) {
	newEntry := &services.Entry{}
	err := decodeAndValidate(r, newEntry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = e.Service.AddEntry(newEntry)
	if err != nil {
		log.Printf("Error adding Entry to DB: %s", err.Error())
	}

	w.WriteHeader(http.StatusCreated)
}

//HandleUpdateEntry handles replacing an entry
func (e *EntryHandler) HandleUpdateEntry(w http.ResponseWriter, r *http.Request) {
	entry := &services.Entry{}
	err := decodeAndValidate(r, entry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if entry.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = e.Service.UpdateEntry(entry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusAccepted)
}

//HandleDeleteEntry handles deleting an entry
func (e *EntryHandler) HandleDeleteEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := e.Service.DeleteEntry(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

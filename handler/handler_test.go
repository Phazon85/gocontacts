package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/phazon85/go_contacts/services"
)

type testDataSerice struct{}

func (tds *testDataSerice) AllEntries() ([]*services.Entry, error) {
	return []*services.Entry{}, nil
}

func (tds *testDataSerice) EntryByID(id string) (*services.Entry, error) {
	return &services.Entry{}, nil
}

func (tds *testDataSerice) AddEntry(entry *services.Entry) error {
	return nil
}

func (tds *testDataSerice) UpdateEntry(entry *services.Entry) error {
	return nil
}

func (tds *testDataSerice) DeleteEntry(id string) error {
	return nil
}

// func (tds *testDataSerice) EntriesToCSV() {
// 	tempFile, err := io.util.TempFile(os.TempDir(), "tmp.*.csv")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return tempFile, nil
// }

// func (tds *testDataSerice) CSVToEntries() {
// 	return []*services.Entry(), nil
// }

func TestHandleGetEntries(t *testing.T) {
	testEntryHandler := NewEntryHandler(&testDataSerice{})

	r := mux.NewRouter()
	r.HandleFunc("/entry", testEntryHandler.HandleGetEntries).Methods("GET")

	req, err := http.NewRequest("GET", "/entry", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code: %d, got : %d", status, http.StatusOK)
	}
}

func TestHandleGetEntryByID(t *testing.T) {
	testEntryhandler := NewEntryHandler(&testDataSerice{})

	r := mux.NewRouter()
	r.HandleFunc("/entry/{id:[0-9]+}", testEntryhandler.HandleGetEntriesByID).Methods("GET")

	req, err := http.NewRequest("GET", "/entry/3", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code: %d, got : %d", status, http.StatusOK)
	}
}

var testEntry = &services.Entry{
	FirstName: "Test",
	LastName:  "Stuff",
	Email:     "Test.Stuff@gmail.com",
	Phone:     "1111111111",
}

func TestHandleAddEntry(t *testing.T) {
	testJSON, err := json.Marshal(testEntry)
	if err != nil {
		t.Error(err)
	}

	testEntryhandler := NewEntryHandler(&testDataSerice{})

	r := mux.NewRouter()
	r.HandleFunc("/entry", testEntryhandler.HandleAddEntry).Methods("POST")

	req, err := http.NewRequest("POST", "/entry", bytes.NewBuffer(testJSON))
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status code: %d, but got %d", status, http.StatusCreated)
	}
}

func TestHandleUpdateEntry(t *testing.T) {
	testJSON, err := json.Marshal(testEntry)
	if err != nil {
		t.Error(err)
	}

	testEntryhandler := NewEntryHandler(&testDataSerice{})

	r := mux.NewRouter()
	r.HandleFunc("/entry", testEntryhandler.HandleUpdateEntry).Methods("PUT")

	req, err := http.NewRequest("PUT", "/entry", bytes.NewBuffer(testJSON))
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("Expected status code: %d, but got %d", status, http.StatusAccepted)
	}
}

func TestHandleDeleteEntry(t *testing.T) {
	testEntryhandler := NewEntryHandler(&testDataSerice{})

	r := mux.NewRouter()
	r.HandleFunc("/entry", testEntryhandler.HandleDeleteEntry).Methods("DELETE")

	req, err := http.NewRequest("DELETE", "/entry/3", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code: %d to equal: %d", status, http.StatusOK)
	}
}

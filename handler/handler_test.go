package handler

import (
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

// func (tds *testDataSerice) AddEntry(entry *services.Entry) (*services.Entry, error) {
// 	return []*services.Entry{}, nil
// }

// func (tds *testDataSerice) UpdateEntry(entry *services.Entry) error {
// 	return nil
// }

// func (tds *testDataSerice) DeleteEntryByID(id string) error {
// 	return nil
// }

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

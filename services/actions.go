package services

import (
	"database/sql"
	"errors"
	"os"
)

var (
	errIDNil        = errors.New("id cannot be nil")
	errInvalidID    = errors.New("ID not found in DB")
	errFirstNameNil = errors.New("FirstName cannot be nil")
	errLastNameNil  = errors.New("LastName cannot be nil")
	errEmailNil     = errors.New("Email cannot be nil")
	errPhoneNil     = errors.New("Phone cannot be nil")
	errNoID         = errors.New("id cannot be nil")
	errEmailExists  = errors.New("Email already exists")
	errNotFound     = errors.New("not found")
)

//Actions implements methods for handling contact entries
type Actions interface {
	AllEntries() ([]*Entry, error)
	EntryByID(id string) (*Entry, error)
	AddEntry(entry *Entry) error
	UpdateEntry(entry *Entry) error
	DeleteEntry(id string) error
	EntriesToCSV() (*os.File, error)
}

//Entry defines a contact entry in the DB
type Entry struct {
	ID        string `json:"ID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Phone     string `json:"Phone"`
}

//PSQLService implements a postgres DB and the actions interface
type PSQLService struct {
	DB *sql.DB
}

//Validation defines testing methods for json
type Validation interface {
	Validate() error
}

//Validate checks fields of struct Entry and returns appropriate errors
func (e Entry) Validate() error {
	if e.FirstName == "" {
		return errFirstNameNil
	}
	if e.LastName == "" {
		return errEmailNil
	}
	if e.Email == "" {
		return errFirstNameNil
	}
	if e.Phone == "" {
		return errPhoneNil
	}
	return nil
}

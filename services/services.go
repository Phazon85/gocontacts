package services

import (
	"database/sql"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

const (
	allEntries         = "SELCET entryid, firstname, lastname, email, phone FROM entries;"
	entryByID          = "SELECT entryid, firstname, lastname, email, phone FROM entries WHERE entryid=%1;"
	createEntry        = "INSERT INTO entries (firstname, lastname, email, phone) VALUES (%2, %3, %4, %5) RETURNING entryid;"
	deleteEntry        = "DELETE FROM entries WHERE entryid = %1"
	updateEntry        = "UPDATE entries SET firstname=%2, lastname=%3, email=%4, phone=%5 WHERE entryid=%1"
	checkIfEmailExists = "SELECT entryid, email FROM entries WHERE email=$1"
)

//InitDB takes in a SQL object for package to use
func InitDB(database *sql.DB) *PSQLService {
	db := database
	return &PSQLService{
		DB: db,
	}
}

//AllEntries returns a list of all contact entries
func (p *PSQLService) AllEntries() ([]*Entry, error) {
	all := []*Entry{}
	rows, err := p.DB.Query(allEntries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		newEntry := &Entry{}
		err = rows.Scan(&newEntry.ID, &newEntry.FirstName, &newEntry.LastName, &newEntry.Email, &newEntry.Phone)
		if err != nil {
			return nil, err
		}
		all = append(all, newEntry)
	}
	return all, err
}

//EntryByID returns a single entry by id
func (p *PSQLService) EntryByID(id string) (*Entry, error) {
	entry := &Entry{}
	if id == "" {
		return nil, errNoID
	}
	row := p.DB.QueryRow(entryByID, id)
	if err := row.Scan(&entry.ID, &entry.FirstName, &entry.LastName, &entry.Email, &entry.Phone); err != nil {
		return nil, err
	}
	return entry, nil
}

func (p *PSQLService) checkIfEmailExists(entry *Entry) error {
	// selector := Entry{
	// 	Email: entry.Email,
	// }

	result := &Entry{}
	row := p.DB.QueryRow(checkIfEmailExists, entry.Email)
	if err := row.Scan(&result.ID, &result.Email); err == nil && result.ID != entry.ID {
		return errEmailExists
	}
	return nil
}

//AddEntry takes an entry and inserts it into the DB generating an ID
func (p *PSQLService) AddEntry(entry *Entry) error {
	entry.Email = strings.ToLower(entry.Email)
	result := p.checkIfEmailExists(entry)
	if result != nil {
		return result
	}

	_, err := p.DB.Exec(createEntry, entry.ID, entry.FirstName, entry.LastName, entry.Email, entry.Phone)
	return err
}

//UpdateEntry will replace an existing entry with new values
func (p *PSQLService) UpdateEntry(entry *Entry) error {
	entry.Email = strings.ToLower(entry.Email)
	result := p.checkIfEmailExists(entry)
	if result != nil {
		return result
	}

	_, err := p.DB.Exec(updateEntry, entry.ID, entry.FirstName, entry.LastName, entry.Email, entry.Phone)
	return err
}

//DeleteEntry will delete a row with given id
func (p *PSQLService) DeleteEntry(id string) error {
	_, err := p.DB.Exec(deleteEntry, id)
	return err
}

//EntriesToCSV will get all entries and write them to a CSV file
func (p *PSQLService) EntriesToCSV() (*os.File, error) {
	entries := []*Entry{}
	rows, err := p.DB.Query(allEntries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		newEntry := &Entry{}
		err = rows.Scan(&newEntry.ID, &newEntry.FirstName, &newEntry.LastName, &newEntry.Email, &newEntry.Phone)
		if err != nil {
			return nil, err
		}
		entries = append(entries, newEntry)
	}

	entriesFile, err := ioutil.TempFile(os.TempDir(), "tmp.*.csv")
	if err != nil {
		return nil, err
	}
	err = gocsv.MarshalFile(&entries, entriesFile)
	if err != nil {
		return nil, err
	}

	return entriesFile, nil
}

package services

import (
	"database/sql"
	"strings"
)

const (
	allEntries         = "SELCET id, firstname, lastname, email, phone FROM entries;"
	entryByID          = "SELECT id, firstname, lastname, email, phone FROM entries WHERE id=%1;"
	createEntry        = "INSERT INTO entries (id, firstname, lastname, email, phone) VALUES (%1, %2, %3, %4, %5) RETURN id;"
	deleteEntry        = "DELETE FROM entries WHERE id = %1"
	updateEntry        = "UPDATE entries SET firstname=%2, lastname=%3, email=%4, phone=%5 WHERE id=%1"
	checkIfEmailExists = "SELECT id, email FROM entries WHERE email=$1"
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
	selector := Entry{
		Email: entry.Email,
	}

	result := &Entry{}
	row := p.DB.QueryRow(checkIfEmailExists, selector.Email)
	if err := row.Scan(&entry.ID, &entry.Email); err == nil && result.ID != entry.ID {
		return errEmailExists
	}
	if err != nil {
		switch err {
			case 
		}
	}

}

//AddEntry takes an entry and inserts it into the DB generating an ID
func (p *PSQLService) AddEntry(entry *Entry) error {
	entry.Email = strings.ToLower(entry.Email)

}

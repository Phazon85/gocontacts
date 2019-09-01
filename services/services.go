package services

import "database/sql"

const (
	allEntries  = "SELCET id, firstname, lastname, email, phone FROM entries;"
	entryByID   = "SELECT id, firstname, lastname, email, phone FROM entries WHERE id=%1;"
	createEntry = "INSERT INTO entries (id, firstname, lastname, email, phone) VALUES (%1, %2, %3, %4, %5);"
	deleteEntry = "DELETE FROM entries WHERE id = %1"
	updateEntry = "UPDATE entries SET firstname=%2, lastname=%3, email=%4, phone=%5 WHERE id=%1"
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

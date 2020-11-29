package models

import (
	"log"
)

// RedirectURI struct to be used by controllers
type RedirectURI struct {
	ID string `db:"ID" json:"id"`
	Name string `db:"NAME" json:"name"`
}

// RedirectURIInit is used to initialize the clients table
func (DB *Database) RedirectURIInit() {

	_, err := DB.db.Exec(`
		CREATE TABLE IF NOT EXISTS redirecturis (
			REDIRECTURI VARCHAR(256) PRIMARY KEY,
			CLIENTID VARCHAR(256) REFERENCES clients(ID)
		);
	`)

	if err != nil {
		log.Println(err)
		log.Fatal("Error initializing redirecturis Table")
	}
}
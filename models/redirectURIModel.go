package models

import (
	"log"
)

// RedirectURI struct to be used by controllers
type RedirectURI struct {
	URI string `db:"REDIRECT_URI" json:"redirectURI"`
	ClientID string `db:"CLIENT_ID" json:"clientID"`
}

// RedirectURIInit is used to initialize the redirecturis table
func (DB *Database) RedirectURIInit() {

	_, err := DB.db.Exec(`
		CREATE TABLE IF NOT EXISTS redirecturis (
			REDIRECT_URI VARCHAR(256) PRIMARY KEY,
			CLIENT_ID VARCHAR(256) NOT NULL REFERENCES clients(ID) 
		);
	`)

	if err != nil {
		log.Println(err)
		log.Fatal("Error initializing redirecturis Table")
	}
}
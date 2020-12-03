package models

import (
	"log"
	"fmt"
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


// CheckRedirectURIForClient checks if the redirectURI is registered for the client
func (DB *Database) CheckRedirectURIForClient(redirectURI string, clientID string) (bool, error) {
	query := fmt.Sprintf("SELECT * FROM redirecturis WHERE REDIRECT_URI='%s' AND CLIENT_ID='%s'",redirectURI,clientID)

	redirectURIs := []RedirectURI{}
	err := DB.db.Select(&redirectURIs,query)

	if (err != nil) || (err == nil && len(redirectURIs) == 0) {
		return false, err
	}

	return true, nil;
}
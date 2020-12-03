package models

import (
	"fmt"
	"log"
)

// Client struct to be used by controllers
type Client struct {
	ID string `db:"ID" json:"id"`
	Name string `db:"NAME" json:"name"`
	Scope string `db:"SCOPE" json:"scope"`
}

// ClientInit is used to initialize the clients table
func (DB *Database) ClientInit() {

	_, err := DB.db.Exec(`
		CREATE TABLE IF NOT EXISTS clients (
			ID VARCHAR(256) PRIMARY KEY,
			NAME VARCHAR(256) NOT NULL,
			SCOPE VARCHAR(256) NOT NULL
		);
	`)

	if err != nil {
		log.Println(err)
		log.Fatal("Error initializing clients Table")
	}
}

// GetClientByClientID is a function that takes in clientID and returns back client
func (DB *Database) GetClientByClientID(clientID string) (Client, error) {
	query := fmt.Sprintf("SELECT * FROM clients WHERE id='%s'",clientID)

	existingClients := []Client{}
	err := DB.db.Select(&existingClients, query)

	if (err != nil) || (err == nil && len(existingClients) == 0) {
		return Client{}, err
	}

	return existingClients[0], nil
}

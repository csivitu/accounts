package models

import (
	"log"
)

// Client struct to be used by controllers
type Client struct {
	ID string `db:"ID" json:"id"`
	Name string `db:"NAME" json:"name"`
}

// ClientInit is used to initialize the clients table
func (DB *Database) ClientInit() {

	_, err := DB.db.Exec(`
		CREATE TABLE IF NOT EXISTS clients (
			ID VARCHAR(256) PRIMARY KEY,
			NAME VARCHAR(256) NOT NULL
		);
	`)

	if err != nil {
		log.Println(err)
		log.Fatal("Error initializing clients Table")
	}
}
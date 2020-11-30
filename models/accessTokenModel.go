package models

import (
	"log"
	"time"
)

// AccessToken struct to be used by controllers
type AccessToken struct {
	Token string `db:"TOKEN" json:"token"`
	Scope string `db:"SCOPE" json:"scope"`
	ExpiresAt time.Time `db:"EXPIRES_AT" json:"expiresAt"`
}

// AcessTokenInit is used to initialize the clients table
func (DB *Database) AcessTokenInit() {

	_, err := DB.db.Exec(`
		CREATE TABLE IF NOT EXISTS accesstokens (
			TOKEN VARCHAR(256) PRIMARY KEY,
			SCOPE VARCHAR(256) NOT NULL,
			EXPIRES_AT DATE NOT NULL
		);
	`)

	if err != nil {
		log.Println(err)
		log.Fatal("Error initializing clients Table")
	}
}
package models

import (
	"fmt"
	"log"

	"github.com/csivitu/accounts/config/dbconfig"

	"github.com/jmoiron/sqlx"
)

// User struct to be used my controllers
type User struct {
	Email    string `db:"EMAIL" json:"email"`
	Username string `db:"USERNAME" json:"username"`
	Password string `db:"PASSWORD" json:"password"`
}

var db *sqlx.DB

func init() {

	db = dbconfig.DB

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			EMAIL         VARCHAR(256) PRIMARY KEY,
			USERNAME      VARCHAR(128) UNIQUE NOT NULL,
			PASSWORD	  VARCHAR(256) NOT NULL
		);
	`)

	if err != nil {
		log.Println(err)
		log.Fatal("Error initializing Users Table")
	}

}

// CreateUser takes a User struct to create user and returns error if any
func CreateUser(user *User) error {
	_, err := db.NamedExec("INSERT INTO users VALUES (:EMAIL, :USERNAME, :PASSWORD)", &user)
	return err
}

// GetUserByUsername is a function that takes in username and returns back user
func GetUserByUsername(username string) (User, error) {

	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s'", username)

	existingUsers := []User{}
	err := db.Select(&existingUsers, query)



	if (err != nil) || (err == nil && len(existingUsers) == 0) {
		return User{},err
	}

	return existingUsers[0],nil
}

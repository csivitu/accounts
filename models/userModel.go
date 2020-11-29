package models

import (
	"fmt"
	"log"
)

// User struct to be used my controllers
type User struct {
	Email    string `db:"EMAIL" json:"email"`
	Username string `db:"USERNAME" json:"username"`
	Password string `db:"PASSWORD" json:"password"`
	Mobile string `db:"MOBILE" json:"mobile"`
	RegNo string `db:"REGNO" json:"regno"`
	Gender string `db:"GENDER" json:"gender"`
}

// UserInit is used to initialize the users table
func (DB *Database) UserInit() {

	_, err := DB.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			EMAIL         VARCHAR(256) PRIMARY KEY,
			USERNAME      VARCHAR(128) UNIQUE NOT NULL,
			PASSWORD	  VARCHAR(256) NOT NULL,
			MOBILE        INT NOT NULL,
			REGNO         VARCHAR(9) UNIQUE,
			GENDER  	  CHAR(1) NOT NULL CHECK(GENDER IN ('M','F'))
		);
	`)

	if err != nil {
		log.Println(err)
		log.Fatal("Error initializing Users Table")
	}

}

// CreateUser takes a User struct to create user and returns error if any
func (DB *Database) CreateUser(user *User) error {
	_, err := DB.db.NamedExec("INSERT INTO users VALUES (:EMAIL, :USERNAME, :PASSWORD)", &user)
	return err
}

// GetUserByUsername is a function that takes in username and returns back user
func (DB *Database) GetUserByUsername(username string) (User, error) {

	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s'", username)

	existingUsers := []User{}
	err := DB.db.Select(&existingUsers, query)

	if (err != nil) || (err == nil && len(existingUsers) == 0) {
		return User{}, err
	}

	return existingUsers[0], nil
}

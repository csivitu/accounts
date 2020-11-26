package models

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
)

var (
	dbUser string
	dbPass string
	dbIP   string
	dbPort int
	dbName string
)

// Database is a wrapper around the SQL database
// to allow addition of methods
type Database struct {
	db *sqlx.DB
}

func init() {
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")

	if dbUser == "" || dbPass == "" {
		log.Fatal("Environment variables DB_USER and DB_PASS were not set.")
	}

	dbIP = os.Getenv("DB_IP")
	dbPortStr := os.Getenv("DB_PORT")

	if dbIP == "" {
		log.Println("DB_IP was not set, using 127.0.0.1.")
		dbIP = "127.0.0.1"
	}

	if dbPortStr == "" {
		log.Println("DB_PORT was not set, using 3306.")
		dbPort = 3306
	} else {
		var err error
		dbPort, err = strconv.Atoi(dbPortStr)

		if err != nil {
			log.Fatal("Invalid DB_PORT.")
		}
	}

	dbName = "csiaccounts"

	db := dbConn(dbUser, dbPass, dbIP, dbPort, "")

	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		log.Println(err)
		log.Fatal("Error creating database csiaccounts")
	}

	db.Close()
}

// Init function is ued to initialize all the tables inside the database
func Init(DB *Database) {

	DB.UserInit()
}

// Init is used to initialize the SQL Database
// func Init() error {
	

// 	defer DB.Close()

// 	_, err = DB.db.Exec("DROP TABLE IF EXISTS users")
// 	if err != nil {
// 		return err
// 	}
	
	// _, err = DB.db.Exec(`
	// 	CREATE TABLE users (
	// 		ID            INT NOT NULL PRIMARY KEY,
	// 		EMAIL         VARCHAR(256) UNIQUE NOT NULL,
	// 		USERNAME      VARCHAR(128) UNIQUE NOT NULL,
	// 		PASSWORD	  VARCHAR(256) NOT NULL
	// 	);
	// `)
	// if err != nil {
	// 	return err
	// }

// 	log.Println("Database initialized, tables created.")
// 	return nil
// }

// New returns an instance of Database
// Use this when you don't want to clear the DB
func New() *Database {
	db := dbConn(dbUser, dbPass, dbIP, dbPort, dbName)

	return &Database{
		db: db,
	}
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/csivitu/accounts/models"
	"github.com/csivitu/accounts/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func handleRequests(DB *models.Database) {
	router := mux.NewRouter()
	routers.Init(router, DB)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Println("PORT was not set, using 10000")
		port = ":10000"
	}

	log.Fatal(http.ListenAndServe(port, router))
}

func main() {
	DB := models.New()
	models.Init(DB)
	handleRequests(DB)
}

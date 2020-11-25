package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/csivitu/accounts/config/routerconfig"
	"github.com/csivitu/accounts/routers"
	_ "github.com/go-sql-driver/mysql"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage")
	fmt.Println("Endpoint hit: homePage")
}

func handleRequests() {

	routers.UserHandler(routerconfig.Router)

	log.Fatal(http.ListenAndServe(":10000", routerconfig.Router))
}

func main() {
	handleRequests()
}

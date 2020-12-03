package routers

import (
	"github.com/csivitu/accounts/controllers"
	"github.com/gorilla/mux"
)

func OAuthHandler(controller *controllers.Controller, router *mux.Router) {
	router.HandleFunc("/authorize", controller.AuthorizeClient).Methods("GET")
	// router.HandleFunc("/authorize", controller.Login).Methods("POST")
}

package routers

import (
	"github.com/csivitu/accounts/controllers"
	"github.com/gorilla/mux"
)

func UserHandler(controller *controllers.Controller, router *mux.Router) {
	router.HandleFunc("/user/signup", controller.Signup).Methods("POST")
	router.HandleFunc("/user/login", controller.Login).Methods("POST")
}

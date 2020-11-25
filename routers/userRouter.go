package routers

import (
	"github.com/csivitu/accounts/controllers"
	"github.com/gorilla/mux"
)

func UserHandler(router *mux.Router) {
	router.HandleFunc("/user/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
}

package routers

import (
	"github.com/csivitu/accounts/controllers"
	"github.com/csivitu/accounts/models"
	"github.com/gorilla/mux"
)

// Init initializes all the route handlers
func Init(router *mux.Router, DB *models.Database) {

	c := controllers.New(DB)

	UserHandler(c, router)
	OAuthHandler(c, router)
	
}
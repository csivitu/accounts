package controllers

import (
	"github.com/csivitu/accounts/models"
)

// Controller struct wrapper around all controller functions
type Controller struct {
	DB *models.Database
}

// New returns a new Controller struct
func New(db *models.Database) *Controller {
	return &Controller{
		DB: db,
	}
}
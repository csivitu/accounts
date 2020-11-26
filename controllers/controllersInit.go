package controllers

import (
	"github.com/csivitu/accounts/models"
)

type Controller struct {
	DB *models.Database
}

func New(db *models.Database) *Controller {
	return &Controller{
		DB: db,
	}
}
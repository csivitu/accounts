package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/csivitu/accounts/models"
	"github.com/csivitu/accounts/utils"
)

// Signup controller function
func (c *Controller) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		User    models.User `json:"user"`
	}

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		json.NewEncoder(w).Encode(&response{Status: "failure", Message: "Error signing up user"})
		return
	}
	user.Password = hashedPassword

	err = c.DB.CreateUser(&user)

	if err != nil {
		json.NewEncoder(w).Encode(&response{Status: "failure", Message: "Error signing up user"})
		return
	}

	json.NewEncoder(w).Encode(response{Status: "success", Message: "User created successfully", User: user})

}

//Login controller function
func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var lgnDta loginData
	json.NewDecoder(r.Body).Decode(&lgnDta)

	existingUser, err := c.DB.GetUserByUsername(lgnDta.Username)

	type response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		User    models.User `json:"user"`
	}

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(&response{
			Status:  "failure",
			Message: "unable to fetch data from database",
		})
		return
	}

	if existingUser == (models.User{}) || utils.CheckPasswordHash(lgnDta.Password, existingUser.Password) != nil {
		json.NewEncoder(w).Encode(&response{
			Status:  "failure",
			Message: "invalid login credentials",
		})
	} else {
		json.NewEncoder(w).Encode(&response{
			Status:  "success",
			Message: "user exists",
			User:    existingUser,
		})
		
	}

}

package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/csivitu/accounts/utils"

	"github.com/csivitu/accounts/models"
)

// AuthorizeClient is used to authorize client for using ouath services
func (c *Controller) AuthorizeClient(w http.ResponseWriter, r *http.Request) {
	scope := r.FormValue("scope")
	responseType := r.FormValue("response_type")
	clientID := r.FormValue("client_id")
	redirectURI := r.FormValue("redirect_uri")
	state := r.FormValue("state")


	if(scope == "" || responseType == "" || clientID == "" || redirectURI == "") {
		utils.AuthResponseError(w,r,"invalid_request","missing query values",state, redirectURI)
	}

	// Check if response type is code
	if(responseType != "code") {
		utils.AuthResponseError(w,r,"invalid_request","invalid responseType",state, redirectURI)
	}
	
	// Check if clientID is present
	existingClient, err := c.DB.GetClientByClientID(clientID)
	if err != nil {
		utils.AuthResponseError(w,r,"server_error","error fetching client",state, redirectURI)
	}
	if(existingClient == (models.Client{})) {
		utils.AuthResponseError(w,r,"acess_denied","unregistered client",state, redirectURI)
	}

	// Check if scope is present in the database
	scopes := strings.Split(existingClient.Scope, " ")
	isScopeIncluded := utils.CheckScope(scopes, scope)
	if(!isScopeIncluded) {
		utils.AuthResponseError(w,r,"acess_denied","unregistered scope",state, redirectURI)
	}


	// Check if redirectURI is present in database
	checkRedirectURI, err := c.DB.CheckRedirectURIForClient(redirectURI, clientID)
	if err != nil {
		utils.AuthResponseError(w,r,"server_error","error fetching client",state, redirectURI)
	}
	if !checkRedirectURI {
		utils.AuthResponseError(w,r,"acess_denied","unregistered redirect uri",state, redirectURI)
	}

	// On correct validation redirect to accounts page
	redirect := fmt.Sprintf("https://accounts.csivit.com/user/signup?event=%s",existingClient.Name)
	http.Redirect(w,r,redirect,http.StatusSeeOther);


}

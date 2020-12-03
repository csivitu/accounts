package utils

import (
    "fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword function to hash a password
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// CheckPasswordHash function to compared hashed passwords
func CheckPasswordHash(password, hash string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err
}

// CheckScope is used to check if teh scope is present
func CheckScope(scopes []string, scope string) bool {
    for _,e := range scopes {
        if e == scope {
            return true;
        }
    }
    return false;
}

func AuthResponseError(w http.ResponseWriter, r *http.Request, errorType string, errorDescription string, state string, redirectURI string) {
    redirect := fmt.Sprintf("%s?error=%s&error_description=%s&state=%s",redirectURI,errorType,errorDescription,state)
    http.Redirect(w,r,redirect,http.StatusSeeOther);
}
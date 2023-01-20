package utils

import (
	"errors"
	"log"
	"net/http"
	"web-project/models"

	"golang.org/x/crypto/bcrypt"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// HashPassword hashes the given password and returns the hashed password and error
func HashPassword(pass string) (string, error) {
	if len(pass) == 0 {
		return "", errors.New("password cannot be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hash), err
}

// ValidatePassword compares given passwor with hashed password
func ValidatePassword(givenPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(givenPass), []byte(pass)) == nil
}

// SendRequest sends a request to the given url and returns the response and error
func SendRequest(url models.Url) (*models.Request, error) {
	resp, err := http.Get(url.Address)
	req := new(models.Request)
	req.UrlId = url.Id
	if err != nil {
		return req, err
	}
	req.Result = resp.StatusCode
	return req, nil
}

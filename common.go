package team

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

//error variables
var (
	ErrorAuth       = errors.New("auth enabled not initialised")
	ErrorPermission = errors.New("permissions enabled not initialised")
)

// HashingPassword pass the arguments password it will return the bcrypt hashed password
func hashingPassword(pass string) string {

	passbyte, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {

		panic(err)

	}

	return string(passbyte)
}

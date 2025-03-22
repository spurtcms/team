package team

import (
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// error variables
var (
	ErrorAuth       = errors.New("auth enabled not initialised")
	ErrorPermission = errors.New("permissions enabled not initialised")
	ErrorValidation = errors.New("given some values is empty")
	TenantId        = os.Getenv("Tenant_ID")
)

// HashingPassword pass the arguments password it will return the bcrypt hashed password
func hashingPassword(pass string) string {

	passbyte, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {

		panic(err)

	}

	return string(passbyte)
}

func AuthandPermission(team *Teams) error {

	//check auth enable if enabled not use auth pkg otherwise it will return error
	if team.AuthEnable && !team.Auth.AuthFlg {

		return ErrorAuth
	}
	//check permission enable if enabled not use team-role pkg otherwise it will return error
	if team.PermissionEnable && !team.Auth.PermissionFlg {

		return ErrorPermission

	}

	return nil
}

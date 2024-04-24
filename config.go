package team

import (
	"github.com/spurtcms/auth"
	"gorm.io/gorm"
)

type Config struct {
	DB               *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	Authenticate     auth.Authentication
	Auth             *auth.Auth
}

type Teams struct {
	DB               *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	Authenticate     auth.Authentication
	Auth             *auth.Auth
}

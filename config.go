package team

import (
	"github.com/spurtcms/auth"
	role "github.com/spurtcms/team-roles"
	"gorm.io/gorm"
)

type Config struct {
	DB               *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	Authenticate     auth.Authentication
	PermissionConf   *role.PermissionConfig
	Auth             *auth.Auth
}

type Teams struct {
	DB               *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	Authenticate     auth.Authentication
	PermissionConf   *role.PermissionConfig
	Auth             *auth.Auth
}

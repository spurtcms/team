package team

import (
	"github.com/spurtcms/auth"
	"gorm.io/gorm"
)

type Type string

const ( //for permission check
	Postgres Type = "postgres"
	Mysql    Type = "mysql"
)

type Config struct {
	DB               *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	DataBaseType     Type
	Auth             *auth.Auth
}

type Teams struct {
	DB               *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	Authenticate     auth.Authentication
	Auth             *auth.Auth
}

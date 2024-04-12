package team

import (
	"fmt"
	"testing"

	"github.com/spurtcms/auth"
	role "github.com/spurtcms/team-roles"
	"gorm.io/gorm"
)

var SecretKey = "Secret123"

func TestTeamList(t *testing.T) {

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  SecretKey,
	})

	token, _ := Auth.CreateToken()

	Permission := role.RoleSetup(role.Config{
		RoleId: 1,
		DB:     &gorm.DB{},
	})

	_, err := Permission.IsGranted("roles", role.CRUD)

	Config := Config{
		DB:               &gorm.DB{},
		AuthEnable:       true,
		PermissionEnable: true,
		Authenticate:     auth.Authentication{Token: token, SecretKey: SecretKey},
		PermissionConf:   Permission,
		Auth:             Auth,
	}

	team := TeamSetup(Config)

	teamuser, count, err := team.ListUser(10, 0, Filters{})

	if err != nil {

		panic(err)
	}

	fmt.Println(teamuser, count)
}

func TestCreateTeam(t *testing.T) {

}

func TestUpdateTeam(t *testing.T) {

}

func TestDeleteteam(t *testing.T) {

}

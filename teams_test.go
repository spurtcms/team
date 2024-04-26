package team

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBSetup() (*gorm.DB, error) {

	dbConfig := map[string]string{
		"username": "postgres",
		"password": "123",
		"host":     "localhost",
		"port":     "5432",
		"dbname":   "spurtcms",
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=" + dbConfig["username"] + " password=" + dbConfig["password"] +
			" dbname=" + dbConfig["dbname"] + " host=" + dbConfig["host"] +
			" port=" + dbConfig["port"] + " sslmode=disable TimeZone=Asia/Kolkata",
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err != nil {
		return nil, err
	}

	return db, nil
}

var SecretKey = "Secret123"

func TestTeamList(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	_, err := Auth.IsGranted("teams", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Authenticate:     auth.Authentication{Token: token, SecretKey: SecretKey},
		Auth:             Auth,
	})

	//list users
	teamuser, count, err := team.ListUser(10, 0, Filters{})

	if err != nil {

		panic(err)
	}

	fmt.Println(teamuser, count)

}

func TestCreateTeam(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Team", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Authenticate:     auth.Authentication{Token: token, SecretKey: SecretKey},
		Auth:             Auth,
	})

	if permisison {

		_, terr := team.CreateUser(TeamCreate{FirstName: "demo", RoleId: 1, Email: "demo@gmail.com"}) // TeamCreate we have multiple fields for creating user details

		if terr != nil {

			log.Println(terr)
		}
	} else {

		log.Println("permissions enabled not initialised")

	}

}

func TestUpdateTeam(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Team", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Authenticate:     auth.Authentication{Token: token, SecretKey: SecretKey},
		Auth:             Auth,
	})
	if permisison {

		team.UpdateUser(TeamCreate{FirstName: "admin", RoleId: 1, Email: "demo@gmail.com"}, 2)

	} else {

		log.Println("permissions enabled not initialised")

	}
}

func TestDeleteteam(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Team", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Authenticate:     auth.Authentication{Token: token, SecretKey: SecretKey},
		Auth:             Auth,
	})
	if permisison {

		team.DeleteUser(2)

	} else {

		log.Println("permissions enabled not initialised")

	}
}

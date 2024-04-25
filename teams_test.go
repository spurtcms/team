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
		"password": "root",
		"host":     "localhost",
		"port":     "5432",
		"dbname":   "spurt-cms-apr3",
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
		RoleId: 1,
		RoleName: "",
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
		RoleId: 1,
		RoleName: "",
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()


	Auth.VerifyToken(token, SecretKey)


	_, err := Auth.IsGranted("teams", auth.CRUD)

	if err != nil {

		panic(err)
	}

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Authenticate:     auth.Authentication{Token: token, SecretKey: SecretKey},
		Auth:             Auth,
	})

	_, terr := team.CreateUser(TeamCreate{FirstName: "demo", RoleId: 1, Email: "demo@gmail.com"}) // TeamCreate we have multiple fields for creating user details

	

	if terr != nil {

		log.Println(terr)
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
		RoleId: 1,
		RoleName: "",
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()


	Auth.VerifyToken(token, SecretKey)


    Auth.IsGranted("teams", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Authenticate:     auth.Authentication{Token: token, SecretKey: SecretKey},
		Auth:             Auth,
	})

	team.UpdateUser(TeamCreate{FirstName: "admin", RoleId: 1, Email: "demo2@gmail.com"}, 2)
}

func TestDeleteteam(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId: 1,
		RoleName: "",
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)


	Auth.IsGranted("teams", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Authenticate:     auth.Authentication{Token: token, SecretKey: SecretKey},
		Auth:             Auth,
	})

	team.DeleteUser(2)
}
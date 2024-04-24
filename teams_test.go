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
		"password": "****",
		"host":     "localhost",
		"port":     "5432",
		"dbname":   "demo",
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

	Auth := auth.AuthSetup(auth.Config{
		ExpiryTime: 2,
		SecretKey:  SecretKey,
		RoleId:     1,
		DB:         db,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token,SecretKey)

	_, err := Auth.IsGranted("teams", auth.CRUD)

	team := TeamSetup(Config{
		DB:              db,
		AuthEnable:       true,
		PermissionEnable: true,
		Authenticate:     auth.Authentication{Token: token, SecretKey: SecretKey},
		Auth:             Auth,
	})

	//list users
	teamuser, count, err := team.ListUser(10, 0, Filters{})

	//create user
	team.CreateUser(TeamCreate{FirstName: "demo", RoleId: 1, Email: "demo@gmail.com"}) // TeamCreate we have multiple fields for creating user details

	//update user
	team.UpdateUser(TeamCreate{FirstName: "demo1", RoleId: 2, Email: "demo1@gmail.com"}, 1)

	//delete user
	team.DeleteUser(1)

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

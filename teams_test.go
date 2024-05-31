package team

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Db connection
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

// TeamList function test
func TestTeamList(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
		RoleName:   "",
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Team", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	//list users
	if permisison {

		teamuser, count, _ := team.ListUser(10, 0, Filters{})

		fmt.Println(teamuser, count)

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// Create team function test
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

// Update team function test
func TestUpdateTeam(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
		RoleName:   "",
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Team", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		team.UpdateUser(TeamCreate{FirstName: "admin", RoleId: 1, Email: "demo@gmail.com"}, 2)

	} else {

		log.Println("permissions enabled not initialised")

	}
}

// Delete team function test
func TestDeleteteam(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
		RoleName:   "",
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Team", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		team.DeleteUser([]int{}, 2, 1)

	} else {

		log.Println("permissions enabled not initialised")

	}
}

// Checkeamil function test
func TestCheckemail(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
		RoleName:   "",
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Team", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		_, chk, _ := team.CheckEmail("demo2@gmail.com", 2)

		log.Println("chk", chk)

	} else {

		log.Println("permissions enabled not initialised")

	}

}

func TestCheckNumber(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
		RoleName:   "",
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Team", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		chk, _ := team.CheckNumber("9900887500", 2)

		log.Println("chk", chk)

	} else {

		log.Println("permissions enabled not initialised")

	}

}
func TestCheckUserValidation(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
		RoleName:   "",
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Team", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		email, user, mobile, _ := team.CheckUserValidation("9900887501", "demo@gmail.com", "demo", 2)

		log.Println("chk", email, user, mobile)

	} else {

		log.Println("permissions enabled not initialised")

	}

}
func TestCheckPasswordwithOld(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
		RoleName:   "",
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Team", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		chk, _ := team.CheckPasswordwithOld(2, "Admin@123")

		log.Println("chk", chk)

	} else {

		log.Println("permissions enabled not initialised")

	}

}

func TestLastLoginActivity(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
		RoleName:   "",
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Team", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		chk := team.LastLoginActivity(2)

		log.Println("chk", chk)

	} else {

		log.Println("permissions enabled not initialised")

	}

}
func TestCheckRoleUsed(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     2,
		RoleName:   "",
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Team", auth.CRUD)

	team := TeamSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})
	if permisison {

		chk, _ := team.CheckRoleUsed(3)

		log.Println("chk", chk)

	} else {

		log.Println("permissions enabled not initialised")

	}

}

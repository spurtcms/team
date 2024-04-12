package team

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Channelsetup used to initialie channel configuration
func TeamSetup(config Config) *Teams {

	MigrationTables(config.DB)

	return &Teams{
		DB:               config.DB,
		AuthEnable:       config.AuthEnable,
		PermissionEnable: config.PermissionEnable,
		Authenticate:     config.Authenticate,
		PermissionConf:   config.PermissionConf,
		Auth:             config.Auth,
	}

}

// get the all list users
func (team *Teams) ListUser(limit, offset int, filter Filters) (tbluserr []tbluser, totoaluser int64, err error) {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return []tbluser{}, 0, AuthError
	}

	UserList, _, terr := tm.GetUsersList(offset, limit, filter, false, team.DB)

	if terr != nil {

		return []tbluser{}, 0, terr
	}

	var userlists []tbluser

	for _, val := range UserList {

		var first = val.FirstName

		var last = val.LastName

		var firstn = strings.ToUpper(first[:1])

		var lastn string

		if val.LastName != "" {

			lastn = strings.ToUpper(last[:1])
		}

		var Name = firstn + lastn

		val.NameString = Name

		userlists = append(userlists, val)

	}

	_, usercount, _ := tm.GetUsersList(0, 0, filter, false, team.DB)

	return userlists, usercount, nil

}

// CreateUser create for your admin login.
func (team *Teams) CreateUser(teamcreate TeamCreate) (createuser tbluser, terr error) {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return tbluser{}, AuthError
	}

	password := teamcreate.Password

	uvuid := (uuid.New()).String()

	hash_pass := hashingPassword(password)

	var user tbluser

	user.Uuid = uvuid

	user.RoleId = teamcreate.RoleId

	user.FirstName = teamcreate.FirstName

	user.LastName = teamcreate.LastName

	user.Email = teamcreate.Email

	user.Username = teamcreate.Username

	user.Password = hash_pass

	user.MobileNo = teamcreate.MobileNo

	user.IsActive = teamcreate.IsActive

	user.DataAccess = teamcreate.DataAccess

	user.ProfileImage = teamcreate.ProfileImage

	user.ProfileImagePath = teamcreate.ProfileImagePath

	user.DefaultLanguageId = 1

	user.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	user.CreatedBy = teamcreate.CreatedBy

	newuser, err := tm.CreateUser(&user, team.DB)

	if err != nil {

		return tbluser{}, err
	}

	return newuser, nil
}

// update user.
func (team *Teams) UpdateUser(teamcreate TeamCreate, userid int) (createuser tbluser, terr error) {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return tbluser{}, AuthError
	}

	user_id := userid

	password := teamcreate.Password

	var user tbluser

	if password != "" {

		hash_pass := hashingPassword(password)

		user.Password = hash_pass
	}

	user.Id = user_id

	user.RoleId = teamcreate.RoleId

	user.FirstName = teamcreate.FirstName

	user.LastName = teamcreate.LastName

	user.Email = teamcreate.Email

	user.Username = teamcreate.Username

	user.MobileNo = teamcreate.MobileNo

	user.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	user.ModifiedBy = user_id

	user.IsActive = teamcreate.IsActive

	user.DataAccess = teamcreate.DataAccess

	user.ProfileImage = teamcreate.ProfileImage

	user.ProfileImagePath = teamcreate.ProfileImagePath

	User, err := tm.UpdateUser(&user, team.DB)

	if err != nil {

		return tbluser{}, err
	}

	return User, nil

}

// delete user.
func (team *Teams) DeleteUser(userid int) error {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return AuthError
	}
	var user tbluser

	user.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	user.DeletedBy = userid

	user.IsDeleted = 1

	user.Id = userid

	err := tm.DeleteUser(&user, team.DB)

	if err != nil {

		return err
	}

	return nil
}

// check email is already exists in your database
func (team *Teams) CheckEmail(Email string, userid int) (users TblUser, checl bool, errr error) {

	var user TblUser

	err := tm.CheckEmail(&user, Email, userid, team.DB)

	if err != nil {

		return TblUser{}, false, err
	}

	return user, true, nil
}

// check mobile
func (team *Teams) CheckNumber(mobile string, userid int) (bool, error) {

	var user TblUser

	err := tm.CheckNumber(&user, mobile, userid, team.DB)

	if err != nil {

		return false, err
	}

	return true, nil
}

// Check username,email,number exsits or not validation
func (team *Teams) CheckUserValidation(mobile string, email string, username string, userid int) (emaill bool, users bool, mobiles bool, err error) {

	var user TblUser

	err1 := tm.CheckValidation(&user, email, username, mobile, userid, team.DB)

	if err1 != nil {

		return false, false, false, err1
	}

	return true, true, true, nil
}

/*check new password with old password*/
/*if it's return false it does not match to the old password*/
/*or return true it does match to the old password*/
func (team *Teams) CheckPasswordwithOld(userid int, password string) (bool, error) {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return false, AuthError
	}

	var user TblUser

	err := tm.GetUserDetailsTeam(&user, userid, team.DB)

	if err != nil {

		return false, err
	}

	passerr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if passerr == bcrypt.ErrMismatchedHashAndPassword {

		return false, nil

	}

	return true, nil
}

/*Logout Last Active*/
func (team *Teams) LastLoginActivity(userid int) (err error) {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return AuthError
	}

	Lerr := tm.Lastlogin(userid, time.Now(), team.DB)

	if Lerr != nil {

		return err
	}

	return nil
}

// Check role already used or not
func (team *Teams) CheckRoleUsed(roleid int) (bool, error) {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return false, AuthError
	}

	var tbluser TblUser

	err := tm.CheckRoleUsed(&tbluser, roleid, team.DB)

	if err != nil {

		return false, err
	}

	return true, nil

}

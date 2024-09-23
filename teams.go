package team

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"time"

	"fmt"

	"github.com/google/uuid"
	"github.com/spurtcms/team/migration"
	"golang.org/x/crypto/bcrypt"
)

// Channelsetup used to initialie channel configuration
func TeamSetup(config Config) *Teams {

	migration.AutoMigration(config.DB, config.DataBaseType)

	return &Teams{
		DB:               config.DB,
		AuthEnable:       config.AuthEnable,
		PermissionEnable: config.PermissionEnable,
		Auth:             config.Auth,
	}

}

// get the all list users
func (team *Teams) ListUser(limit, offset int, filter Filters, Tenantid int) (tbluserr []TblUser, totoaluser int64, err error) {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return []TblUser{}, 0, AuthError
	}

	tm.Userid = team.Userid
	tm.Dataaccess = team.Dataaccess

	UserList, _, terr := tm.GetUsersList(offset, limit, filter, false, true, team.DB, Tenantid)

	if terr != nil {

		return []TblUser{}, 0, terr
	}

	var userlists []TblUser

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

	_, usercount, _ := tm.GetUsersList(0, 0, filter, false, true, team.DB, Tenantid)

	return userlists, usercount, nil

}

// CreateUser create for your admin login.
func (team *Teams) CreateUser(teamcreate TeamCreate) (createuser TblUser, UserId int, terr error) {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return TblUser{}, 0, AuthError
	}

	password := teamcreate.Password
	uvuid := (uuid.New()).String()
	hash_pass := hashingPassword(password)

	var user TblUser

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
	user.StorageType = teamcreate.StorageType
	user.TenantId = teamcreate.TenantId
	user.S3FolderName = teamcreate.S3FolderPath

	fmt.Println("user:", teamcreate.TenantId)

	newuser, err := tm.CreateUser(&user, team.DB)

	if err != nil {

		return TblUser{}, 0, err
	}

	Userid, err := tm.GetUserByRole(teamcreate.RoleId, teamcreate.MobileNo, team.DB)

	if err != nil {

		return TblUser{}, 0, err
	}

	return newuser, Userid, nil
}

func (team *Teams) CreateTenantid(user TblMstrTenant) (int, error) {
	id, err := tm.CreateTenantid(&user, team.DB)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (team *Teams) UpdateTenantId(UserId int, Tenantid int) {
	tm.UpdateTenantId(UserId, Tenantid, team.DB)
}

func (team *Teams) UpdateS3FolderName(tenantId, userId int, s3FolderPath string) error {
	err := tm.UpdateS3FolderName(tenantId, userId, s3FolderPath, team.DB)
	if err != nil {
		return err
	}

	return nil
}

func (team *Teams) GetTenantDetails(tenantId int) (tenantDetails TblUser, err error) {

	tenantDetails, err = tm.GetTenantDetails(tenantId, team.DB)
	if err != nil {
		return TblUser{}, err
	}

	return tenantDetails, nil
}

func (team *Teams) UpdateImageDetails(userId int, imageName, imagepath string) (err error) {

	err = tm.UpdateImageDetails(userId, imageName, imagepath, team.DB)
	if err != nil {
		return err
	}

	return nil
}

func (team *Teams) CreateTenantApiToken(UserId int, tenantId int) (ApiToken string, err error) {
	ApiToken, err = GenerateTenantApiToken(64)
	if err != nil {
		return "", err
	}
	currentTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	tokenDetails := TblGraphqlSettings{
		TokenName:   "Default Token",
		Description: "Default token",
		Duration:    "Unlimited",
		CreatedOn:   currentTime,
		Token:       ApiToken,
		IsDefault:   1,
		TenantId:    tenantId}
	switch {
	case UserId != 0:
		tokenDetails.CreatedBy = UserId
	}
	err = tm.CreateTenantApiToken(team.DB, &tokenDetails)
	if err != nil {
		return "", err
	}
	return ApiToken, nil
}

func GenerateTenantApiToken(length int) (string, error) {
	b := make([]byte, length)               // Create a slice to hold 32 bytes of random data
	if _, err := rand.Read(b); err != nil { // Fill the slice with random data and handle any errors
		return "", err // Return an empty string and the error if something went wrong
	}
	return base64.URLEncoding.EncodeToString(b), nil // Encode the random bytes to a URL-safe base64 string
}

// update user.
func (team *Teams) UpdateUser(teamcreate TeamCreate, userid int, tenantid int) (createuser TblUser, terr error) {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return TblUser{}, AuthError
	}

	user_id := userid

	password := teamcreate.Password

	var user TblUser

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
	user.StorageType = teamcreate.StorageType

	tm.Userid = team.Userid
	tm.Dataaccess = team.Dataaccess
	User, err := tm.UpdateUser(&user, team.DB, tenantid)

	if err != nil {

		return TblUser{}, err
	}

	return User, nil

}

// delete user.
func (team *Teams) DeleteUser(usersIds []int, userid int, deletedby int, tenantid int) error {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return AuthError
	}

	var user TblUser

	user.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	user.DeletedBy = deletedby
	user.IsDeleted = 1

	err := tm.DeleteMultipleUser(&user, usersIds, userid, team.DB, tenantid)

	if err != nil {

		return err
	}

	return nil
}

// check email is already exists in your database
func (team *Teams) CheckEmail(Email string, userid int, tenantid int) (users TblUser, checl bool, errr error) {

	var user TblUser

	err := tm.CheckEmail(&user, Email, userid, team.DB, tenantid)

	if err != nil {

		return TblUser{}, false, err
	}

	return user, true, nil
}

// check mobile
func (team *Teams) CheckNumber(mobile string, userid int, tenantid int) (bool, error) {

	var user TblUser

	err := tm.CheckNumber(&user, mobile, userid, team.DB, tenantid)

	if err != nil {

		return false, err
	}

	return true, nil
}

// Check username,email,number exsits or not validation
func (team *Teams) CheckUserValidation(mobile string, email string, username string, userid int, tenantid int) (emaill bool, users bool, mobiles bool, err error) {

	var user TblUser

	err1 := tm.CheckValidation(&user, email, username, mobile, userid, team.DB, tenantid)

	if err1 != nil {

		return false, false, false, err1
	}

	return true, true, true, nil
}

/*check new password with old password*/
/*if it's return false it does not match to the old password*/
/*or return true it does match to the old password*/
func (team *Teams) CheckPasswordwithOld(userid int, password string, tenantid int) (bool, error) {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return false, AuthError
	}

	var user TblUser

	err := tm.GetUserDetailsTeam(&user, userid, team.DB, tenantid)

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
func (team *Teams) LastLoginActivity(userid int, tenantid int) (err error) {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return AuthError
	}

	Lerr := tm.Lastlogin(userid, time.Now(), team.DB, tenantid)

	if Lerr != nil {

		return err
	}

	return nil
}

// Check role already used or not
func (team *Teams) CheckRoleUsed(roleid int, tenantid int) (bool, error) {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return false, AuthError
	}

	var TblUser TblUser

	err := tm.CheckRoleUsed(&TblUser, roleid, team.DB, tenantid)

	if err != nil {

		return false, err
	}

	return true, nil

}

// get team by id
func (team *Teams) GetUserById(Userid int, Userids []int) (tbluser TblUser, users []TblUser, err error) {

	//check if auth or permission enabled
	if autherr := AuthandPermission(team); autherr != nil {

		return TblUser{}, []TblUser{}, autherr
	}

	user, users, err := tm.GetUserById(Userid, Userids, team.DB)

	if err != nil {

		return TblUser{}, []TblUser{}, err
	}

	return user, users, nil

}

func (team *Teams) UserDetails(inputs Team) (UserDetails TblUser, err error) {

	//check if auth or permission enabled
	if autherr := AuthandPermission(team); autherr != nil {

		return TblUser{}, autherr
	}

	tm.Userid = team.Userid
	tm.Dataaccess = team.Dataaccess

	err = tm.GetUserDetails(team.DB, inputs, &UserDetails)

	if err != nil {

		return TblUser{}, err
	}

	return UserDetails, nil
}

// check username
func (team *Teams) CheckUsername(username string, userid int, tenantid int) (bool, error) {

	var user TblUser

	err := tm.CheckUsername(&user, username, userid, team.DB, tenantid)

	if err != nil {

		return false, err
	}

	return true, nil
}

// change user Access for multiple user
func (team *Teams) ChangeAccess(userIds []int, modifiedby int, status int, tenantid int) error {

	var user TblUser

	user.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	user.ModifiedBy = modifiedby
	user.DataAccess = status

	err := tm.ChangeAccess(&user, userIds, team.DB, tenantid)

	if err != nil {

		return err
	}

	return nil
}

// change active status
func (team *Teams) ChangeActiveStatus(userId int, activeStatus int, modifiedby int, tenantid int) (bool, error) {

	var userStatus TblUser

	userStatus.ModifiedBy = modifiedby
	userStatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	userStatus.IsActive = activeStatus

	err := tm.ChangeActiveUser(&userStatus, userId, team.DB, tenantid)

	if err != nil {

		return false, err

	}
	return true, nil

}

// change active Status for multiple users
func (team *Teams) SelectedUserStatusChange(userIds []int, activeStatus int, modifiedby int, tenantid int) error {

	var userActiveStatus TblUser

	userActiveStatus.ModifiedBy = modifiedby
	userActiveStatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	userActiveStatus.IsActive = activeStatus

	err := tm.SelectedUserStatusChange(&userActiveStatus, userIds, team.DB, tenantid)

	if err != nil {

		return err
	}

	return nil

}

// Dashboard usercount function
func (team *Teams) DashboardUserCount(tenantid int) (totalcount int, lasttendayscount int, err error) {

	if autherr := AuthandPermission(team); autherr != nil {
		return 0, 0, autherr
	}

	allusercount, err := tm.UserCount(team.DB, tenantid)
	if err != nil {
		return 0, 0, err
	}

	lusercount, err := tm.NewuserCount(team.DB, tenantid)
	if err != nil {
		return 0, 0, err
	}

	return int(allusercount), int(lusercount), nil
}

func (team *Teams) ChangeYourPassword(password string, userid int, tenantid int) (success bool, err error) {

	if autherr := AuthandPermission(team); autherr != nil {
		return false, autherr
	}

	var tbluser TblUser
	tbluser.Id = userid
	hash_pass := hashingPassword(password)
	tbluser.Password = hash_pass
	tbluser.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	tbluser.ModifiedBy = 1

	cerr := tm.ChangePasswordById(&tbluser, team.DB, tenantid)
	if cerr != nil {
		return false, cerr
	}

	return true, nil
}

func (team *Teams) GetAdminRoleUsers(roleid []int, tenantid int) (userlist []TblUser, err error) {

	if autherr := AuthandPermission(team); autherr != nil {
		return []TblUser{}, autherr
	}
	userslist, _ := tm.GetAdminRoleUsers(roleid, team.DB, tenantid)

	return userslist, nil

}

func (team *Teams) UpdateMyUser(userupdate TeamCreate, userid int, tenantid int) error {

	if autherr := AuthandPermission(team); autherr != nil {
		return autherr
	}

	if userupdate.FirstName == "" || userupdate.Email == "" || userupdate.Username == "" || userupdate.MobileNo == "" {
		return ErrorValidation
	}

	password := userupdate.Password

	var user TblUser
	if password != "" {
		hash_pass := hashingPassword(password)
		user.Password = hash_pass
	}

	user.Id = userid
	user.FirstName = userupdate.FirstName
	user.LastName = userupdate.LastName
	user.Email = userupdate.Email
	user.Username = userupdate.Username
	user.MobileNo = userupdate.MobileNo
	user.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	user.ModifiedBy = userid
	user.DataAccess = userupdate.DataAccess
	user.ProfileImage = userupdate.ProfileImage
	user.ProfileImagePath = userupdate.ProfileImagePath
	user.StorageType = userupdate.StorageType

	err := tm.UpdateMyuser(&user, team.DB, tenantid)

	if err != nil {
		return err
	}

	return nil
}

// change active Status for multiple users
func (team *Teams) ChangeStatusForTenants(userIds []int, activeStatus int, modifiedby int) error {

	var userActiveStatus TblUser

	userActiveStatus.ModifiedBy = modifiedby
	userActiveStatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	userActiveStatus.IsActive = activeStatus

	err := tm.ChangeStatusForTenants(&userActiveStatus, userIds, team.DB)

	if err != nil {

		return err
	}

	return nil

}

// delete only tenant users
func (team *Teams) DeleteTenantusers(usersIds []int, userid int, deletedby int) error {

	if AuthError := AuthandPermission(team); AuthError != nil {

		return AuthError
	}

	var user TblUser

	user.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	user.DeletedBy = deletedby
	user.IsDeleted = 1

	err := tm.DeleteTenantusers(&user, usersIds, userid, team.DB)

	if err != nil {

		return err
	}

	return nil
}

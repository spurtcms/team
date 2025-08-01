package team

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TblUser struct {
	Id                   int       `gorm:"column:id"`
	Uuid                 string    `gorm:"column:uuid"`
	FirstName            string    `gorm:"column:first_name"`
	LastName             string    `gorm:"column:last_name"`
	RoleId               int       `gorm:"column:role_id"`
	Email                string    `gorm:"column:email"`
	Username             string    `gorm:"column:username"`
	Password             string    `gorm:"column:password"`
	MobileNo             string    `gorm:"column:mobile_no"`
	IsActive             int       `gorm:"column:is_active"`
	ProfileImage         string    `gorm:"column:profile_image"`
	ProfileImagePath     string    `gorm:"column:profile_image_path"`
	StorageType          string    `gorm:"column:storage_type"`
	DataAccess           int       `gorm:"column:data_access"`
	CreatedOn            time.Time `gorm:"column:created_on"`
	CreatedBy            int       `gorm:"column:created_by"`
	ModifiedOn           time.Time `gorm:"column:modified_on;DEFAULT:NULL"`
	ModifiedBy           int       `gorm:"column:modified_by;DEFAULT:NULL"`
	LastLogin            time.Time `gorm:"column:last_login;DEFAULT:NULL"`
	IsDeleted            int       `gorm:"column:is_deleted"`
	DeletedOn            time.Time `gorm:"column:deleted_on;DEFAULT:NULL"`
	DeletedBy            int       `gorm:"column:deleted_by;DEFAULT:NULL"`
	ModuleName           string    `gorm:"-"`
	RouteName            string    `gorm:"-:migration;<-:false"`
	DisplayName          string    `gorm:"-:migration;<-:false"`
	Description          string    `gorm:"-"`
	ModuleId             int       `gorm:"-:migration;<-:false"`
	PermissionId         int       `gorm:"-"`
	FullAccessPermission int       `gorm:"-:migration;<-:false"`
	RoleName             string    `gorm:"-:migration;<-:false"`
	DefaultLanguageId    int       `gorm:"column:default_language_id"`
	NameString           string    `gorm:"-"`
	TenantId             string
	Otp                  int       `gorm:"column:otp"`
	OtpExpiry            time.Time `gorm:"column:otp_expiry;DEFAULT:NULL"`
	NameLength           int       `gorm:"-:migration;<-:false"`
	LimitedLengthName    string    `gorm:"-:migration;<-:false"`
	S3FolderName         string    `gorm:"column:s3_folder_name"`
	Subdomain            string
	GoTemplateDefault    int `gorm:"column:go_template_default"`
	UsageMode            string
	ChannelId            int
	Country              string `gorm:"column:country"`
}

type TblMstrTenant struct {
	Id            int       `gorm:"primaryKey;auto_increment;type:serial"`
	TenantId      string    `gorm:"type:character varying"`
	S3StoragePath string    `gorm:"type:character varying"`
	DeletedOn     time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy     int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted     int       `gorm:"type:integer;DEFAULT:0"`
}

type Filters struct {
	Keyword   string
	Category  string
	Status    string
	FromDate  string
	ToDate    string
	FirstName string
}

type Team struct {
	Id         int
	Limit      int
	Offset     int
	Keyword    string
	FirstName  string
	TenantId   string
	IsActive   bool
	CreateOnly bool
	Count      bool
	Role       bool
	EmailId    string
}

type TeamCreate struct {
	FirstName         string
	LastName          string
	RoleId            int
	Email             string
	Username          string
	Password          string
	IsActive          int
	DataAccess        int
	MobileNo          string
	ProfileImage      string
	ProfileImagePath  string
	CreatedBy         int
	StorageType       string
	TenantId          string
	S3FolderPath      string
	DefaultLanguageId int
}

type TeamModel struct {
	Dataaccess int
	Userid     int
}

var tm TeamModel

// get the list of users
func (t TeamModel) GetUsersList(offset, limit int, filter Filters, flag bool, createonly bool, DB *gorm.DB, tenantid string) (users []TblUser, count int64, err error) {

	var Total_users int64

	query := DB.Table("tbl_users").Select("tbl_users.id,tbl_users.uuid,tbl_users.role_id,tbl_users.first_name,tbl_users.last_name,tbl_users.email,tbl_users.password,tbl_users.username,tbl_users.mobile_no,tbl_users.profile_image,tbl_users.profile_image_path,tbl_users.created_on,tbl_users.created_by,tbl_users.modified_on,tbl_users.modified_by,tbl_users.is_active,tbl_users.is_deleted,tbl_users.deleted_on,tbl_users.deleted_by,tbl_users.data_access,tbl_roles.name as role_name,tbl_users.storage_type").Joins("inner join tbl_roles on tbl_users.role_id = tbl_roles.id")

	if tenantid == "" {
		query = query.Where("tbl_users.is_deleted=? and tbl_users.created_by = ?", 0, t.Userid)
	} else {
		query = query.Where("tbl_users.is_deleted=? and tbl_users.tenant_id=?", 0, tenantid)
	}
	if createonly && t.Dataaccess == 1 {

		query = query.Where("tbl_users.created_by = ?", t.Userid)
	}

	if filter.Keyword != "" {

		query = query.Where("(LOWER(TRIM(tbl_users.first_name)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%").
			Or("LOWER(TRIM(tbl_users.last_name)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%").
			Or("LOWER(TRIM(tbl_roles.name)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%").
			Or("LOWER(TRIM(tbl_users.username)) LIKE LOWER(TRIM(?)))", "%"+filter.Keyword+"%")

	}

	if filter.FirstName != "" {

		query = query.Debug().Where("LOWER(TRIM(tbl_users.first_name)) LIKE LOWER(TRIM(?))"+" OR LOWER(TRIM(tbl_users.last_name)) LIKE LOWER(TRIM(?))", "%"+filter.FirstName+"%", "%"+filter.FirstName+"%")

	}

	if flag {

		query.Order("id desc").Find(&users)

		return users, 0, nil

	}

	if limit != 0 && !flag {

		query.Offset(offset).Limit(limit).Order("id desc").Find(&users)

		return users, 0, nil

	}

	query.Find(&users).Count(&Total_users)

	if err1 := query.Error; err1 != nil {

		return []TblUser{}, 0, err1

	}

	return []TblUser{}, Total_users, nil

}

// This func will help to create a user in your database
func (t TeamModel) CreateUser(user *TblUser, DB *gorm.DB) (team TblUser, terr error) {

	if err := DB.Debug().Table("tbl_users").Create(&user).Error; err != nil {

		return TblUser{}, err

	}

	return *user, nil
}

// function to get user id by role id,mobile number  from db
func (t TeamModel) GetUserByRole(RoleId int, MobileNo string, DB *gorm.DB) (id int, err error) {

	var UserId TblUser

	if err := DB.Table("tbl_users").Where("role_id=? and mobile_no=?", RoleId, MobileNo).First(&UserId).Error; err != nil {

		return 0, err
	}

	return UserId.Id, err
}

func (t TeamModel) CreateTenantid(user *TblMstrTenant, DB *gorm.DB) (int, error) {
	result := DB.Table("tbl_mstr_tenants").Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.Id, nil
}

func (t TeamModel) UpdateTenantId(UserId int, Tenantid string, DB *gorm.DB) error {

	result := DB.Table("tbl_users").Where("id = ?", UserId).Update("tenant_id", Tenantid)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t TeamModel) UpdateS3FolderName(tenantId, userId int, s3FolderPath string, DB *gorm.DB) error {

	result := DB.Table("tbl_users").Where("id = ?", userId).Update("s3_folder_name", s3FolderPath)
	if result.Error != nil {
		return result.Error
	}

	result = DB.Table("tbl_mstr_tenants").Where("id = ?", tenantId).Update("s3_storage_path", s3FolderPath)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t TeamModel) UpdateImageDetails(userId int, imageName, imagePath string, DB *gorm.DB) error {

	result := DB.Table("tbl_users").Where("id = ?", userId).UpdateColumns(map[string]interface{}{"profile_image": imageName, "profile_image_path": imagePath})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t TeamModel) GetTenantDetails(tenantId string, DB *gorm.DB) (tenantDetails TblUser, err error) {
	if tenantId == "" {
		result := DB.Debug().Table("tbl_users").Where("tenant_id = ?  and is_deleted = 0 and is_active = 1", tenantId).First(&tenantDetails)
		if result.Error != nil {
			return TblUser{}, result.Error
		}

	} else {
		result := DB.Debug().Table("tbl_users").Where("tenant_id = ? and is_deleted = 0 and is_active = 1", tenantId).First(&tenantDetails)
		if result.Error != nil {
			return TblUser{}, result.Error
		}

	}

	return tenantDetails, nil
}

// update user
func (t TeamModel) UpdateUser(user *TblUser, DB *gorm.DB, tenantid string) (team TblUser, terr error) {
	fmt.Println("user idfff", t.Userid)
	query := DB.Table("tbl_users")
	if tenantid == "" {
		query = query.Where("id=? and  created_by = ?", user.Id, t.Userid)
	} else {
		query = query.Where("id=? and tenant_id=?", user.Id, tenantid)

	}
	if user.ProfileImage == "" || user.Password == "" {

		if user.Password == "" && user.ProfileImage == "" {

			query = query.Omit("password", "profile_image", "profile_image_path").UpdateColumns(map[string]interface{}{"first_name": user.FirstName, "last_name": user.LastName, "role_id": user.RoleId, "email": user.Email, "username": user.Username, "mobile_no": user.MobileNo, "is_active": user.IsActive, "modified_on": user.ModifiedOn, "modified_by": user.ModifiedBy, "data_access": user.DataAccess})

		} else if user.ProfileImage == "" {

			query = query.Omit("profile_image", "profile_image_path").UpdateColumns(map[string]interface{}{"first_name": user.FirstName, "last_name": user.LastName, "role_id": user.RoleId, "email": user.Email, "username": user.Username, "mobile_no": user.MobileNo, "is_active": user.IsActive, "modified_on": user.ModifiedOn, "modified_by": user.ModifiedBy, "data_access": user.DataAccess, "password": user.Password})

		} else if user.Password == "" {

			query = query.Omit("password").UpdateColumns(map[string]interface{}{"first_name": user.FirstName, "last_name": user.LastName, "role_id": user.RoleId, "email": user.Email, "username": user.Username, "mobile_no": user.MobileNo, "is_active": user.IsActive, "modified_on": user.ModifiedOn, "modified_by": user.ModifiedBy, "profile_image": user.ProfileImage, "profile_image_path": user.ProfileImagePath, "data_access": user.DataAccess})
		}

		if err := query.Error; err != nil {

			return TblUser{}, err
		}

	} else {

		if err := query.UpdateColumns(map[string]interface{}{"first_name": user.FirstName, "last_name": user.LastName, "role_id": user.RoleId, "email": user.Email, "username": user.Username, "mobile_no": user.MobileNo, "is_active": user.IsActive, "modified_on": user.ModifiedOn, "modified_by": user.ModifiedBy, "profile_image": user.ProfileImage, "profile_image_path": user.ProfileImagePath, "data_access": user.DataAccess, "password": user.Password}).Error; err != nil {

			return TblUser{}, err
		}
	}
	return *user, nil
}

// delete team user
func (t TeamModel) DeleteUser(user *TblUser, DB *gorm.DB, tenantid string) error {

	if err := DB.Model(&TblUser{}).Where("id=? and tenant_id=?", user.Id, tenantid).Updates(TblUser{IsDeleted: user.IsDeleted, DeletedOn: user.DeletedOn, DeletedBy: user.DeletedBy}).Error; err != nil {

		return err

	}
	return nil
}

// user last login update
func (t TeamModel) Lastlogin(id int, log_time time.Time, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_users").Where("id=? and tenant_id=?", id, tenantid).Update("last_login", log_time).Error; err != nil {

		return err
	}
	return nil

}

func (t TeamModel) GetUserDetailsTeam(user *TblUser, id int, DB *gorm.DB, tenantid string) error {

	if err := DB.Where("id=? and tenant_id=?", id, tenantid).First(&user).Error; err != nil {

		return err
	}
	return nil
}

func (t TeamModel) CheckValidation(user *TblUser, email, username, mobile string, userid int, DB *gorm.DB, tenantid string) error {
	if userid == 0 {
		if err := DB.Table("tbl_users").Where("mobile_no = ? or LOWER(TRIM(email))=LOWER(TRIM(?)) or username = ?   and is_deleted=0 and tenant_id=?", mobile, email, username, tenantid).First(&user).Error; err != nil {

			return err
		}
	} else {
		if err := DB.Table("tbl_users").Where("mobile_no = ? or LOWER(TRIM(email))=LOWER(TRIM(?)) or username = ? and id not in (?) and is_deleted=0 and tenant_id=?", mobile, email, username, userid, tenantid).First(&user).Error; err != nil {

			return err
		}

	}

	return nil
}

func (t TeamModel) CheckEmail(user *TblUser, email string, userid int, DB *gorm.DB, tenantid string) error {

	if userid == 0 {
		if err := DB.Table("tbl_users").Where("LOWER(TRIM(email))=LOWER(TRIM(?)) and is_deleted = 0", email).First(&user).Error; err != nil {

			return err
		}
	} else {
		if err := DB.Table("tbl_users").Where("LOWER(TRIM(email))=LOWER(TRIM(?)) and id not in(?) and is_deleted= 0", email, userid).First(&user).Error; err != nil {

			return err
		}
	}

	return nil
}

func (t TeamModel) CheckNumber(user *TblUser, mobile string, userid int, DB *gorm.DB, tenantid string) error {

	if userid == 0 {
		if err := DB.Table("tbl_users").Where("mobile_no = ? and is_deleted=0", mobile).First(&user).Error; err != nil {

			return err
		}
	} else {
		if err := DB.Table("tbl_users").Where("mobile_no = ? and id not in (?) and is_deleted=0", mobile, userid).First(&user).Error; err != nil {

			return err
		}

	}

	return nil
}

// Rolechekc
func (t TeamModel) CheckRoleUsed(user *TblUser, roleid int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_users").Where("role_id=? and is_deleted =0 and tenant_id=?", roleid, tenantid).First(user).Error; err != nil {
		return err
	}
	return nil

}

// getuserbyid
func (t TeamModel) GetUserById(id int, ids []int, DB *gorm.DB) (user TblUser, users []TblUser, err error) {

	if id != 0 {

		if err := DB.Table("tbl_users").Where("id=?", id).First(&user).Error; err != nil {

			return TblUser{}, []TblUser{}, err
		}

	} else if len(ids) > 0 {

		if err := DB.Table("tbl_users").Where("id in (?)", ids).Find(&users).Error; err != nil {

			return TblUser{}, []TblUser{}, err
		}
	}

	return user, users, nil
}

func (team TeamModel) GetUserDetails(DB *gorm.DB, inputs Team, user *TblUser) error {

	query := DB.Debug().Model(TblUser{}).Where("is_deleted = 0")

	if inputs.CreateOnly && team.Dataaccess == 1 {

		query = query.Where("tbl_users.created_by = ?", team.Userid)
	}

	if inputs.Keyword != "" {

		query = query.Where("(LOWER(TRIM(tbl_users.first_name)) LIKE LOWER(TRIM(?))", "%"+inputs.Keyword+"%").
			Or("LOWER(TRIM(tbl_users.last_name)) LIKE LOWER(TRIM(?))", "%"+inputs.Keyword+"%").
			Or("LOWER(TRIM(tbl_roles.name)) LIKE LOWER(TRIM(?))", "%"+inputs.Keyword+"%").
			Or("LOWER(TRIM(tbl_users.username)) LIKE LOWER(TRIM(?)))", "%"+inputs.Keyword+"%")

	}

	if inputs.FirstName != "" {

		query = query.Debug().Where("LOWER(TRIM(tbl_users.first_name)) LIKE LOWER(TRIM(?))"+" OR LOWER(TRIM(tbl_users.last_name)) LIKE LOWER(TRIM(?))", "%"+inputs.FirstName+"%", "%"+inputs.FirstName+"%")

	}

	if inputs.TenantId != "" {

		query = query.Where("tbl_users.tenant_id=?", inputs.TenantId)
	}

	if inputs.EmailId != "" {

		query = query.Where("email = ? and is_deleted = 0", inputs.EmailId)
	}

	if inputs.Role {

		query.Preload("Role", "is_deleted=?", 0)
	}

	if err := query.First(&user).Error; err != nil {

		return err
	}

	return nil
}

// check username
func (t TeamModel) CheckUsername(user *TblUser, username string, userid int, DB *gorm.DB, tenantid string) error {

	if userid == 0 {

		if err := DB.Table("tbl_users").Where("username = ? and is_deleted=0 and tenant_id=?", username, tenantid).First(&user).Error; err != nil {

			return err
		}

	} else {

		if err := DB.Table("tbl_users").Where("username = ? and id not in (?) and is_deleted=0 and tenant_id=?", username, userid, tenantid).First(&user).Error; err != nil {

			return err
		}

	}

	return nil
}

// change selected user access
func (t TeamModel) ChangeAccess(user *TblUser, userIds []int, DB *gorm.DB, tenantid string) error {

	result := DB.Debug().Model(&user).Where("id IN (?) and tenant_id=?", userIds, tenantid).UpdateColumns(map[string]interface{}{"modified_on": user.ModifiedOn, "modified_by": user.ModifiedBy, "data_access": user.DataAccess})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// change active Status for multiple users

func (t TeamModel) SelectedUserStatusChange(userStatus *TblUser, userIds []int, DB *gorm.DB, tenantid string) error {

	if tenantid == "" {
		if err := DB.Debug().Table("tbl_users").Where("id in (?) ", userIds).UpdateColumns(map[string]interface{}{"is_active": userStatus.IsActive, "modified_by": userStatus.ModifiedBy, "modified_on": userStatus.ModifiedOn}).Error; err != nil {

			return err
		}
	} else {
		if err := DB.Debug().Table("tbl_users").Where("id in (?) and tenant_id=?", userIds, tenantid).UpdateColumns(map[string]interface{}{"is_active": userStatus.IsActive, "modified_by": userStatus.ModifiedBy, "modified_on": userStatus.ModifiedOn}).Error; err != nil {

			return err
		}
	}

	return nil

}

// Delete Multiple User Functionality

func (t TeamModel) DeleteMultipleUser(user *TblUser, usersIds []int, userid int, DB *gorm.DB, tenantid string) error {

	if tenantid == "" {
		if userid != 0 {
			if err := DB.Model(&TblUser{}).Where("id=? ", userid).Updates(TblUser{IsDeleted: user.IsDeleted, DeletedOn: user.DeletedOn, DeletedBy: user.DeletedBy}).Error; err != nil {

				return err

			}
			return nil
		} else {
			if err := DB.Model(&TblUser{}).Where("id IN (?)", usersIds).Updates(map[string]interface{}{"is_deleted": user.IsDeleted, "deleted_on": user.DeletedOn, "deleted_by": user.DeletedBy}).Error; err != nil {

				return err

			}
		}

	} else {
		if userid != 0 {
			if err := DB.Model(&TblUser{}).Where("id=? and tenant_id=?", userid, tenantid).Updates(TblUser{IsDeleted: user.IsDeleted, DeletedOn: user.DeletedOn, DeletedBy: user.DeletedBy}).Error; err != nil {

				return err

			}
			return nil
		} else {
			if err := DB.Model(&TblUser{}).Where("id IN (?) and tenant_id=?", usersIds, tenantid).Updates(map[string]interface{}{"is_deleted": user.IsDeleted, "deleted_on": user.DeletedOn, "deleted_by": user.DeletedBy}).Error; err != nil {

				return err

			}
		}

	}
	return nil
}

// change active status
func (t TeamModel) ChangeActiveUser(user *TblUser, userId int, DB *gorm.DB, tenantid string) error {

	if tenantid == "" {
		result := DB.Debug().Model(&user).Where("id = ? ", userId).UpdateColumns(map[string]interface{}{"modified_on": user.ModifiedOn, "modified_by": user.ModifiedBy, "is_active": user.IsActive})
		if result.Error != nil {
			return result.Error
		}
	} else {
		result := DB.Debug().Model(&user).Where("id = ? and tenant_id=?", userId, tenantid).UpdateColumns(map[string]interface{}{"modified_on": user.ModifiedOn, "modified_by": user.ModifiedBy, "is_active": user.IsActive})
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (t TeamModel) NewuserCount(DB *gorm.DB, tenantid string) (count int64, err error) {

	if err := DB.Table("tbl_users").Where("is_deleted = 0 AND created_on >=? and tenant_id=?", time.Now().AddDate(0, 0, -10), tenantid).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil

}

func (t TeamModel) UserCount(DB *gorm.DB, tenantid string) (count int64, err error) {

	if err := DB.Table("tbl_users").Where("is_deleted = 0 and tenant_id=?", tenantid).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil

}

func (t TeamModel) ChangePasswordById(user *TblUser, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_users").Where("id=? and tenant_id=?", user.Id, tenantid).Updates(TblUser{Password: user.Password, ModifiedOn: user.ModifiedOn, ModifiedBy: user.ModifiedBy}).Error; err != nil {

		return err
	}
	return nil
}

func (t TeamModel) GetAdminRoleUsers(roleid []int, DB *gorm.DB, tenantid string) (userlist []TblUser, err error) {
	fmt.Println(roleid, "roleid")

	if err := DB.Debug().Table("tbl_users").Where("role_id in (?) and is_active=1 and is_deleted=0 and tenant_id=?", roleid, tenantid).Find(&userlist).Error; err != nil {

		return []TblUser{}, err
	}
	return userlist, nil

}

func (t TeamModel) UpdateMyuser(user *TblUser, DB *gorm.DB, tenantid string) error {

	query := DB.Debug().Table("tbl_users").Where("id=? and tenant_id=?", user.Id, tenantid)

	if user.ProfileImage == "" || user.Password == "" {

		if user.Password == "" && user.ProfileImage == "" {

			query = query.Omit("password", "profile_image", "profile_image_path").UpdateColumns(map[string]interface{}{"first_name": user.FirstName, "last_name": user.LastName, "email": user.Email, "username": user.Username, "mobile_no": user.MobileNo, "modified_on": user.ModifiedOn, "modified_by": user.ModifiedBy, "data_access": user.DataAccess, "storage_type": user.StorageType})

		} else if user.ProfileImage == "" {

			query = query.Omit("profile_image", "profile_image_path").UpdateColumns(map[string]interface{}{"first_name": user.FirstName, "last_name": user.LastName, "email": user.Email, "username": user.Username, "mobile_no": user.MobileNo, "modified_on": user.ModifiedOn, "modified_by": user.ModifiedBy, "data_access": user.DataAccess, "password": user.Password, "storage_type": user.StorageType})

		} else if user.Password == "" {

			query = query.Omit("password").UpdateColumns(map[string]interface{}{"first_name": user.FirstName, "last_name": user.LastName, "email": user.Email, "username": user.Username, "mobile_no": user.MobileNo, "modified_on": user.ModifiedOn, "modified_by": user.ModifiedBy, "profile_image": user.ProfileImage, "profile_image_path": user.ProfileImagePath, "data_access": user.DataAccess, "storage_type": user.StorageType})
		}

		if err := query.Error; err != nil {

			return err
		}

	} else {

		if err := query.UpdateColumns(map[string]interface{}{"first_name": user.FirstName, "last_name": user.LastName, "email": user.Email, "username": user.Username, "mobile_no": user.MobileNo, "modified_on": user.ModifiedOn, "modified_by": user.ModifiedBy, "profile_image": user.ProfileImage, "profile_image_path": user.ProfileImagePath, "data_access": user.DataAccess, "password": user.Password, "storage_type": user.StorageType}).Error; err != nil {

			return err
		}

	}

	return nil
}

// change active Status for multiple users

func (t TeamModel) ChangeStatusForTenants(userStatus *TblUser, userIds []int, DB *gorm.DB) error {

	if err := DB.Debug().Table("tbl_users").Where("id in (?)", userIds).UpdateColumns(map[string]interface{}{"is_active": userStatus.IsActive, "modified_by": userStatus.ModifiedBy, "modified_on": userStatus.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil

}

// delete only tenant users
func (t TeamModel) DeleteTenantusers(user *TblUser, usersIds []int, userid int, DB *gorm.DB) error {

	if userid != 0 {
		if err := DB.Model(&TblUser{}).Where("id=? ", userid).Updates(TblUser{IsDeleted: user.IsDeleted, DeletedOn: user.DeletedOn, DeletedBy: user.DeletedBy}).Error; err != nil {

			return err

		}
		return nil
	} else {
		if err := DB.Model(&TblUser{}).Where("id IN (?) ", usersIds).Updates(map[string]interface{}{"is_deleted": user.IsDeleted, "deleted_on": user.DeletedOn, "deleted_by": user.DeletedBy}).Error; err != nil {

			return err

		}
	}

	return nil
}

// update goTemplate by id
func (t TeamModel) UpdateGoTemplateById(templateid int, userid int, tenantid string, DB *gorm.DB) error {

	if err := DB.Model(&TblUser{}).Where("id=? and tenant_id=?", userid, tenantid).Updates(TblUser{GoTemplateDefault: templateid}).Error; err != nil {

		return err

	}

	return nil
}

func (t TeamModel) CheckDomainName(subdomain string, userid int, tenantid string, DB *gorm.DB) error {

	var userdet TblUser

	if err := DB.Table("tbl_users").Where("subdomain = ? AND id != ? AND is_deleted = 0 ", subdomain, userid).First(&userdet).Error; err != nil {

		return err
	}

	return nil
}

func (t TeamModel) UpdateUserDetails(userdata map[string]interface{}, userid int, tenantid string, DB *gorm.DB) error {

	if err := DB.Table("tbl_users").Where("id=? and is_deleted = 0 and tenant_id=?", userid, tenantid).UpdateColumns(userdata).Error; err != nil {

		return err
	}

	return nil
}

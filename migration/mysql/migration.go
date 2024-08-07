package mysql

import (
	"time"

	"gorm.io/gorm"
)

type TblRole struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	Name        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	Slug        string    `gorm:"type:varchar(255)"`
	IsActive    int       `gorm:"type:int"`
	IsDeleted   int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:int"`
	TenantId    int       `gorm:"type:int"`
}

type TblUser struct {
	Id                int       `gorm:"primaryKey"`
	Uuid              string    `gorm:"type:varchar(255)"`
	FirstName         string    `gorm:"type:varchar(255)"`
	LastName          string    `gorm:"type:varchar(255)"`
	RoleId            TblRole   `gorm:"type:int;foreignkey:Id"`
	Email             string    `gorm:"type:varchar(255)"`
	Username          string    `gorm:"type:varchar(255)"`
	Password          string    `gorm:"type:varchar(255)"`
	MobileNo          string    `gorm:"type:varchar(255)"`
	IsActive          int       `gorm:"type:int"`
	ProfileImage      string    `gorm:"type:varchar(255)"`
	ProfileImagePath  string    `gorm:"type:varchar(255)"`
	DataAccess        int       `gorm:"type:int"`
	DefaultLanguageId int       `gorm:"type:int"`
	CreatedOn         time.Time `gorm:"type:datetime"`
	CreatedBy         int       `gorm:"type:int"`
	ModifiedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy        int       `gorm:"DEFAULT:NULL;type:int"`
	LastLogin         time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted         int       `gorm:"type:int"`
	DeletedOn         time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy         int       `gorm:"DEFAULT:NULL;type:int"`
	TenantId          int       `gorm:"type:int"`
}

type TblMstrTenant struct {
	Id        int       `gorm:"primaryKey;auto_increment"`
	TenantId  int       `gorm:"type:int"`
	DeletedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:int;DEFAULT:NULL"`
	IsDeleted int       `gorm:"type:int;DEFAULT:0"`
}

func MigrationTables(db *gorm.DB) {

	err := db.AutoMigrate(
		&TblRole{},
		&TblUser{},
		&TblMstrTenant{},
	)

	if err != nil {

		panic(err)
	}

}

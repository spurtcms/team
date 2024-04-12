package team

import (
	"time"

	"gorm.io/gorm"
)

type TblRole struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name        string    `gorm:"type:character varying"`
	Description string    `gorm:"type:character varying"`
	Slug        string    `gorm:"type:character varying"`
	IsActive    int       `gorm:"type:integer"`
	IsDeleted   int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy   int       `gorm:"type:integer"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:integer"`
}

type TblLanguage struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	LanguageName string    `gorm:"type:character varying"`
	LanguageCode string    `gorm:"type:character varying"`
	ImagePath    string    `gorm:"type:character varying"`
	IsStatus     int       `gorm:"type:integer"`
	IsDefault    int       `gorm:"type:integer"`
	JsonPath     string    `gorm:"type:character varying"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy    int       `gorm:"type:integer"`
	ModifiedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy   int       `gorm:"DEFAULT:NULL"`
	DeletedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy    int       `gorm:"DEFAULT:NULL"`
	IsDeleted    int       `gorm:"DEFAULT:0"`
}

type TblUser struct {
	Id                int       `gorm:"primaryKey;type:serial"`
	Uuid              string    `gorm:"type:character varying"`
	FirstName         string    `gorm:"type:character varying"`
	LastName          string    `gorm:"type:character varying"`
	RoleId            int       `gorm:"type:integer"`
	Email             string    `gorm:"type:character varying"`
	Username          string    `gorm:"type:character varying"`
	Password          string    `gorm:"type:character varying"`
	MobileNo          string    `gorm:"type:character varying"`
	IsActive          int       `gorm:"type:integer"`
	ProfileImage      string    `gorm:"type:character varying"`
	ProfileImagePath  string    `gorm:"type:character varying"`
	DataAccess        int       `gorm:"type:integer"`
	DefaultLanguageId int       `gorm:"type:integer"`
	CreatedOn         time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy         int       `gorm:"type:integer"`
	ModifiedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy        int       `gorm:"DEFAULT:NULL;type:integer"`
	LastLogin         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted         int       `gorm:"type:integer"`
	DeletedOn         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy         int       `gorm:"DEFAULT:NULL"`
}

func MigrationTables(db *gorm.DB) {

	err := db.AutoMigrate(
		&TblRole{},
		&TblUser{},
		&TblLanguage{},
	)

	if err != nil {

		panic(err)
	}

}

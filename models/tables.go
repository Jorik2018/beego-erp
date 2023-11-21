package models

import "time"

type UserMasterTable struct {
	UserId      int       `orm:"auto"`
	FirstName   string    `orm:"size(255)"`
	LastName    string    `orm:"size(255)"`
	Email       string    `orm:"size(255)"`
	Password    string    `orm:"size(255)"`
	Mobile      string    `orm:"size(255)"`
	CreatedDate time.Time `orm:"type(datetime)"`
}

type CarsMasterTable struct {
	CarsId      int    `orm:"auto"`
	Name        string `orm:"size(255)"`
	Description string `orm:"size(255)"`
	CreatedBy   int
	UpdatedBy   int
	CreatedDate time.Time `orm:"type(date)"`
	UpdatedDate time.Time `orm:"type(date)"`
	// CreatedBy   *UserMasterTable `orm:"rel(fk);null"`
	// UpdatedBy   *UserMasterTable `orm:"rel(fk);null"`
}

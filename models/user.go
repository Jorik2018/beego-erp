package models

import (
	requestStruct "beego-erp/requstStruct"
	"errors"
	"time"

	"github.com/beego/beego/orm"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	UserList["user_11111"] = &u
}

type User struct {
	Id       string
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func RegisterUser(u requestStruct.InsertUser) (interface{}, error) {
	db := orm.NewOrm()
	res := UserMasterTable{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		Password:    u.Password,
		Mobile:      u.Mobile,
		CreatedDate: time.Now(),
	}
	_, err := db.Insert(&res)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func LoginUser(u requestStruct.LoginUser) int {
	db := orm.NewOrm()

	res := UserMasterTable{
		Email:    u.Email,
		Password: u.Password,
	}
	result := db.Read(&res, "Email", "Password")
	if result != nil {
		return 0
	}
	return 1
}

func LoginUsers(u requestStruct.LoginUser) UserMasterTable {
	db := orm.NewOrm()

	res := UserMasterTable{
		Email:    u.Email,
		Password: u.Password,
	}
	result := db.Read(&res, "Email", "Password")
	if result != nil {
		panic("result")
	}
	return res
}

// func InsertnewBook(data InsertBook) (interface{}, error) {
// 	o := orm.NewOrm()
// 	book := Book{
// 		Name:    data.Name,
// 		Author:  data.Author,
// 		Created: time.Now(),
// 	}

// 	_, err := o.Insert(&book)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return book, nil
// }

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}

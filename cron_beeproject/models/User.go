package models

import "github.com/beego/beego/v2/client/orm"

type User struct {

}

func (u *User) TableName() string {
	return "user"
}

func init()  {
	orm.RegisterModel(new(User))
}
package models

import (
	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*User
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id       string `orm:"column(id);pk"`
	Username string
	Password string
	Sex      int
	Age      int
	Phone    int
}

func GetAllUsers() []User {
	o := orm.NewOrm()
	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable("user")
	var us []User
	_, err := qs.All(&us)
	if err == nil {
	}
	return us
}

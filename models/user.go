package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id           int64  `form:"-"`
	Username     string `form:"Username"`
	Password     string `form:"Password"`
	Sex          int    `form:"Sex"`
	Age          int    `form:"Age"`
	Phone        string `form:"Phone"`
	Email        string `form:"Email"`
	Avatarurl    string `form:"Avatarurl"`    // 头像
	Roleid       int    `form:"Roleid"`       //权限id
	Pid          string `form:"Pid"`          //父级id
	Systemtype   string `form:"Systemtype"`   //系统类型
	Address      string `form:"Address"`      //地址
	Placetypes   string `form:"Placetypes"`   //商家类型
	Generatecode string `form:"Generatecode"` //推荐码
	Sharemoney   string `form:"Sharemoney"`   //分享金额
	Codeurl      string `form:"Codeurl"`      //二维码图片
	Openid       string `form:"Openid"`       // 小程序openid
}

func GetAllUsers(Username string) []User {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	var us []User
	_, err := qs.Filter("Username__icontains", Username).All(&us)
	if err == nil {
	}
	return us
}

func AddUser(user *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	return id, err
}

func DeleteUser(user *User) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Delete(user)
	return num, err
}

func GetUserById(Id int64) (user User, err error) {
	o := orm.NewOrm()
	user = User{Id: Id}
	err = o.Read(&user, "Id")
	return
}

func GetUserByName(Username string) (user User, err error) {
	o := orm.NewOrm()
	user = User{Username: Username}
	err = o.Read(&user, "Username")
	return
}

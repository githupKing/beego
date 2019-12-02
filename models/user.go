package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*User
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
	Roleid       string `form:"Roleid"`       //权限id
	Pid          string `form:"Pid"`          //父级id
	Systemtype   string `form:"Systemtype"`   //系统类型
	Address      string `form:"Address"`      //地址
	Placetypes   string `form:"Placetypes"`   //商家类型
	Generatecode string `form:"Generatecode"` //推荐码
	Sharemoney   string `form:"Sharemoney"`   //分享金额
	Codeurl      string `form:"Codeurl"`      //二维码图片
	Openid       string `form:"Openid"`       // 小程序openid
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

func GetUserById(Id int64) (user User, error error) {
	o := orm.NewOrm()
	user = User{Id: Id}
	err := o.Read(&user, "Id")
	return user, err
}

func GetUserByName(Username string) (user User, error error) {
	o := orm.NewOrm()
	fmt.Println(Username)
	user = User{Username: Username}
	err := o.Read(&user, "Username")
	return user, err
}

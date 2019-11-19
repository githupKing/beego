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
	Id           int64
	Username     string
	Password     string
	Sex          int
	Age          int
	Phone        string
	Email        string
	Avatarurl    string // 头像
	Roleid       string //权限id
	Pid          string //父级id
	Systemtype   string //系统类型
	Address      string //地址
	Placetypes   string //商家类型
	Generatecode string //推荐码
	Sharemoney   string //分享金额
	Codeurl      string //二维码图片
	Openid       string // 小程序openid
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

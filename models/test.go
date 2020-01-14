package models

import (
	"github.com/astaxie/beego/orm"
)

type class struct {
	Id        int64   `form:"-"`
	ClassName string  `form:"ClassName"`
	Student   string  `form:"Student"`
	ClassNum  string  `form:"ClassNum"`
	Baby      []*Baby `json:"baby" orm:"reverse(many)"`
}

type Baby struct {
	Id    int64
	Name  string `json:"name" orm:"size(50)"`
	Class *class `json:"user" orm:"rel(fk);index"`
}

func init() {
	orm.RegisterModel(new(class), new(Baby))
}

func GetAllText() []Baby {
	o := orm.NewOrm()
	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable("Baby")
	var us []Baby
	_, err := qs.All(&us)
	if err == nil {
	}
	return us
}

/*
* @Author: wangyongping
* @Date:   2020-01-12 21:41:42
* @Last Modified by:   Administrator
* @Last Modified time: 2020-01-12 21:43:14
 */
package models

import (
	"github.com/astaxie/beego/orm"
)

type Leavel struct {
	Id         int64  `form:"-"`          //权限id
	RoleName   string `form:"RoleName"`   //权限名称
	RoleLeavel string `form:"RoleLeavel"` //权限级别
}

func init() {
	orm.RegisterModel(new(Leavel))
}

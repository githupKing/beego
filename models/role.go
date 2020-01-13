/*
* @Author: wangyongping
* @Date:   2020-01-12 21:34:25
* @Last Modified by:   Administrator
* @Last Modified time: 2020-01-12 22:01:34
 */
package models

import (
	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id           int64  `form:"-"`            //权限id
	RoleLeavelId int    `form:"RoleLeavelId"` //权限级别id
	RoleType     string `form:RoleType`       //权限类型
	RoleName     string `form:RoleName`
}

func init() {
	orm.RegisterModel(new(Role))
}

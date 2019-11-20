package controllers

import (
	"beego/models"
	"github.com/astaxie/beego"
	"strconv"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *UserController) GetAll() {
	users := models.GetAllUsers()
	this.Data["json"] = users
	this.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [post]
func (this *UserController) PostData() {
	user := models.User{}
	_, error := models.AddUser(&user)
	if error != nil {
		this.Ctx.WriteString("服务器错误")
	}
}

// @Title DeleteData
// @Description delete data
// @Success 200 {object} models.User
// @router / [delete]
func (this *UserController) DeleteData() {
	id, _ := strconv.ParseInt(this.GetString("id"), 10, 64)
	user := models.User{Id: id}
	_, err := models.DeleteUser(&user)
	if err != nil {
		this.Ctx.WriteString("服务器错误")
	} else {
		this.Ctx.WriteString("删除成功")
	}
}

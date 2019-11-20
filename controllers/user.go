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
	user := models.User{Username: "王永平", Password: "12346", Sex: 1, Age: 29, Phone: "13208266337", Email: "357754663@qq.com", Avatarurl: "", Roleid: "1", Pid: "1", Systemtype: "1", Address: "四川成都", Placetypes: "1", Generatecode: "12312", Sharemoney: "23", Codeurl: "1.png", Openid: "123123"}
	_, error := models.AddUser(&user)
	if error != nil {
		this.Ctx.WriteString("服务器错误")
	} else {
		this.Ctx.WriteString("添加成功")
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

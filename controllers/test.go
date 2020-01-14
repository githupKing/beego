package controllers

import (
	"beego/models"
	"github.com/astaxie/beego"
)

// Operations about Users
type TestController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *TestController) GetAll() {
	users := models.GetAllText()
	this.Data["json"] = users
	this.ServeJSON()
}

package controllers

import (
	"beego/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

// Operations about Users
type UserController struct {
	beego.Controller
}
type jsons struct {
	Code       int
	Msg        string
	Error_code int
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
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
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

// @Title findOne
// @Description findOne data
// @Success 200 {object} models.User
// @router /findone [get]
func (this *UserController) FindOne() {
	id, _ := strconv.ParseInt(this.GetString("id"), 10, 64) //强类型转换
	user, err := models.GetUserById(id)
	datas := this.Ctx.Request
	fmt.Println(datas.RemoteAddr)
	if err != nil {
		data := &jsons{100, "暂无数据", 1}
		this.Data["json"] = data
		this.ServeJSON()
	} else {
		this.Data["json"] = user
		this.ServeJSON()
	}
}

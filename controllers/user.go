package controllers

import (
	"beego/models"
	"beego/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
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
	Token      string
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *UserController) GetAll() {
	Username := this.GetString("Username")
	users, err := models.GetAllUsers(Username)
	if err != nil {
		this.Ctx.WriteString("出错了！")
	}
	this.Data["json"] = users
	this.ServeJSON()
}

// @Title GetAll
// @Description get all Users 添加用户
// @Success 200 {object} models.User
// @router / [post]
func (this *UserController) PostData() {
	user := models.User{}
	data := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &user); err != nil { //传入user指针
		this.Ctx.WriteString("出错了！")
	} else {
		password := []byte(user.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost) //密码加密
		user.Password = string(hashedPassword)
		if err != nil {
			panic(err)
		}
		_, error := models.AddUser(&user)

		if error != nil {
			this.Ctx.WriteString("服务器错误")
		} else {
			this.Ctx.WriteString("添加成功")
		}
	}
}

// @Title DeleteData
// @Description delete data 删除
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
// @Description findOne data 根据用户id 查询
// @Success 200 {object} models.User
// @router /findone [get]
func (this *UserController) FindOne() {
	id, _ := strconv.ParseInt(this.GetString("id"), 10, 64) //强类型转换
	user, err := models.GetUserById(id)
	if err != nil {
		data := &jsons{100, "暂无数据", 1, ""}
		this.Data["json"] = data
		this.ServeJSON()
	} else {
		this.Data["json"] = user
		this.ServeJSON()
	}
}

// @Title findOne
// @Description findOne data 根据用户名查询
// @Success 200 {object} models.User
// @router /finduserbyname [get]
func (this *UserController) FindUserByName() {
	Username := this.GetString("Username")
	user, err := models.GetUserByName(Username)
	if err != nil {
		data := &jsons{100, "暂无数据", 1, ""}
		this.Data["json"] = data
		this.ServeJSON()
	} else {
		this.Data["json"] = user
		this.ServeJSON()
	}
}

// @Title login
// @Description login data 登录
// @Success 200 {object} models.User
// @router /login [post]
func (this *UserController) Login() {
	UserName := this.GetString("userName")
	fmt.Println(UserName)
	pwd := this.GetString("pwd")
	if UserName == "" || pwd == "" {
		data := &jsons{100, "用户名密码不能为空", 1, ""}
		this.Data["json"] = data
		this.ServeJSON()
		return
	}
	user, err := models.GetUserByName(UserName)
	password := []byte(pwd)
	if err != nil {
		data := &jsons{100, "账号不存在，请先注册", 1, ""}
		this.Data["json"] = data
		this.ServeJSON()
	} else {
		if comparePasswords(user.Password, password) {
			claims := util.Claims{}
			claims.Uid = user.Id
			token, _ := util.CreateToken(&claims)
			data := &jsons{100, "登录成功", 0, token}
			this.Data["json"] = data
			this.ServeJSON()
		} else {
			data := &jsons{100, "账号或密码错误", 1, ""}
			this.Data["json"] = data
			this.ServeJSON()
		}
	}
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

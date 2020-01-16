// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego/controllers"
	"beego/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type jsons struct {
	Code       int
	Msg        string
	Error_code int
	Token      string
}

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(crossDomain),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),
	)
	beego.AddNamespace(ns)
	var BeforeExecFunc = func(this *context.Context) {
		token := this.Input.Header("Authorization")
		if this.Request.RequestURI != "/v1/user/login" && token == "" {
			data := &jsons{400, "请先登录", 1, ""}
			this.Output.JSON(data, false, false)
		}
		_, status := util.ValidateToken(token)
		if this.Request.RequestURI != "/v1/user/login" && !status {
			this.Output.SetStatus(200)
			data := &jsons{400, "登录超时", 1, ""}
			this.Output.JSON(data, false, false)
		}
	}
	beego.InsertFilter("*", beego.BeforeExec, BeforeExecFunc)
}
func crossDomain(ctx *context.Context) {
	// utils.Display("跨域设置", "成功")
	//管理接口
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
	ctx.Output.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	ctx.Output.Header("Access-Control-Allow-Headers", "*, X-Requested-With,Authorization, X-Prototype-Version, X-CSRF-Token, Content-Type")

	if ctx.Input.Method() == "OPTIONS" {
		ctx.Output.SetStatus(200)
		ctx.Output.JSON("{}", false, false)
	}
}

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
		if token == "" {
			data := &jsons{400, "请先登录", 1, ""}
			this.Output.JSON(data, false, false)
		}
		_, status := util.ValidateToken(token)
		if !status {
			this.Output.SetStatus(200)
			data := &jsons{400, "登录超时", 1, ""}
			this.Output.JSON(data, false, false)
		}
	}
	beego.InsertFilter("*", beego.BeforeExec, BeforeExecFunc)
}

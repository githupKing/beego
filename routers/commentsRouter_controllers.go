package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego/controllers:TestController"] = append(beego.GlobalControllerRouter["beego/controllers:TestController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "PostData",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteData",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "FindOne",
            Router: `/findone`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "FindUserByName",
            Router: `/finduserbyname`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

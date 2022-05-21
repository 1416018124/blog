package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["myblog/controllers:BaseController"] = append(beego.GlobalControllerRouter["myblog/controllers:BaseController"],
		beego.ControllerComments{
			Method:           "Homeget",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myblog/controllers:BaseController"] = append(beego.GlobalControllerRouter["myblog/controllers:BaseController"],
		beego.ControllerComments{
			Method:           "ExitGet",
			Router:           "/exit/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myblog/controllers:RegisterController"] = append(beego.GlobalControllerRouter["myblog/controllers:RegisterController"],
		beego.ControllerComments{
			Method:           "LoginGet",
			Router:           "/login/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myblog/controllers:RegisterController"] = append(beego.GlobalControllerRouter["myblog/controllers:RegisterController"],
		beego.ControllerComments{
			Method:           "LoginPost",
			Router:           "/login/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myblog/controllers:RegisterController"] = append(beego.GlobalControllerRouter["myblog/controllers:RegisterController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/register/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["myblog/controllers:RegisterController"] = append(beego.GlobalControllerRouter["myblog/controllers:RegisterController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/register/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}

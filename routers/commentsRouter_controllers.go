package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["myblog/controllers:RegisterController"] = append(beego.GlobalControllerRouter["myblog/controllers:RegisterController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/register/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myblog/controllers:RegisterController"] = append(beego.GlobalControllerRouter["myblog/controllers:RegisterController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/register/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

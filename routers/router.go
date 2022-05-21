package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"myblog/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.RegisterController{}, &controllers.BaseController{})
}

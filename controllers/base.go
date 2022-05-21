package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	Loginuser interface{}
}

// Prepare 判断是否登录
/*
	这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，
    用户可以重写这个函数实现类似用户验证之类。
*/
func (this *BaseController) Prepare() {
	loginuser := this.GetSession("loginuser")
	if loginuser != nil {
		this.IsLogin = true
	} else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}

// Home 代码部分

// @router / [get]
func (this *BaseController) Homeget() {
	logs.Info("%s Login type:", this.Loginuser, this.IsLogin)
	this.TplName = "home.html"
}

//EXIT 用户登录部分
// @router /exit/ [get]
func (this *BaseController) ExitGet() {
	this.DelSession("loginuser")
	this.Data["isLogin"] = false
	this.Redirect("/", 302)
}

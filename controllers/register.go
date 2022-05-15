package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"myblog/models"
	"myblog/utils"
	"time"
)

type RegisterController struct {
	beego.Controller
}

// @router /register/ [get]
func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

// @router /register/ [post]
func (this *RegisterController) Post() {
	//前端使用表单形式传递
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	//判断密码是否一至，前后端都需判断
	if repassword != password {
		this.send_json(403, "密码不一至", "False")
		return
	}

	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误 (数据库username字段设置唯一即可)
	o := orm.NewOrm()
	user := models.User{Username: username}
	err := o.Read(&user)
	if err != orm.ErrNoRows {
		logs.Info("%s 重复", username)
		this.send_json(200, "用户名已经存在", "nil")
		return
	} else if err == orm.ErrNoRows {
		logs.Info("%s 没有重复", username)
	} else {
		logs.Error("%s 查询异常：", err)
		this.send_json(500, "数据库异常", err)
		return
	}

	//写入数据库
	password = utils.MD5(password)
	u := models.User{0, username, password, 0, time.Time{}}
	err = models.UserInsert(u)
	if err != nil {
		this.send_json(500, "注册失败", "True")
	} else {
		this.send_json(200, "注册成功", "False")
	}
}

func (this *RegisterController) send_json(code int, msg string, data interface{}) {
	type returnDataStruct struct {
		Code int
		Msg  string
		Data interface{}
	}

	returnJson := returnDataStruct{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	ret, err := json.Marshal(returnJson)
	if err != nil {
		logs.Error("json marshal error:", err)
	}
	this.Ctx.ResponseWriter.Write(ret)
	this.StopRun()
}

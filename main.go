package main

import (
	_ "myblog/routers"
	"myblog/utils"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	utils.InitMysql()   //初始化Mysql
	utils.InitLogger()  //调用logger初始化
	utils.InitSession() //InitSession
}

func main() {
	beego.Run()
}

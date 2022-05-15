package utils

import (
	"fmt"
	"time"

	"myblog/models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
)

func InitMysql() {

	fmt.Println("InitMysql....")
	driverName, _ := config.String("driverName")

	//注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接
	user, _ := config.String("mysqluser")
	pwd, _ := config.String("mysqlpwd")
	host, _ := config.String("host")
	port, _ := config.String("port")
	dbname, _ := config.String("dbname")

	maxIdle := 30
	maxConn := 30

	//dbConn := "root:123@tcp(127.0.0.1:3306)/myblog?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	err := orm.RegisterDataBase("default", "mysql", dbConn, orm.MaxIdleConnections(maxIdle), orm.MaxOpenConnections(maxConn))
	orm.DefaultTimeLoc = time.UTC
	if err != nil {
		logs.Info("连接数据库出错")
		return
	}
	logs.Info("连接数据库成功")

	//数据库配置
	orm.RegisterModel(new(models.User))
	orm.RunSyncdb("default", false, true)

}

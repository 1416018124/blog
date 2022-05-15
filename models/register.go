package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	Id         int    `orm:"auto"`
	Username   string `orm:"unique"`
	Password   string
	Status     int       // 0 正常状态， 1删除
	Createtime time.Time `orm:"auto_now_add;type(datetime)"`
}

func UserInsert(user User) error {
	var o orm.Ormer
	o = orm.NewOrm()
	_, err := o.Insert(&user)
	return err
}

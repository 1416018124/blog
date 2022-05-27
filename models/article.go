package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type Article struct {
	Id         int `orm:"auto"`
	Title      string
	Tags       string
	Short      string
	Content    string `orm:"type(text)"`
	Author     string
	Createtime time.Time `orm:"auto_now_add;type(datetime)"`
}

func ArticleInsert(article Article) error {
	var o orm.Ormer
	o = orm.NewOrm()
	_, err := o.Insert(&article)
	return err
}

func FindArticleWithPage(page int) ([]Article, error) {
	//从配置文件中获取每页文章数量
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	return QueryArticleWithPage(page, num)
}

func QueryArticleWithPage(page, num int) ([]Article, error) {
	o := orm.NewOrm()
	var article []Article
	//logs.Info(num, page*num)
	_, err := o.QueryTable("article").Limit(num, page*num).All(&article)
	if err != nil {
		logs.Error("查询数据库异常:", err)
	}
	//logs.Info(article)
	return article, err
}

func QueryArticleWithId(id int) Article {
	o := orm.NewOrm()
	var article Article
	//logs.Info(num, page*num)
	err := o.QueryTable("article").Filter("id", id).One(&article)
	if err != nil {
		logs.Error("查询数据库异常:", err)
	}
	return article
}

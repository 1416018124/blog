package controllers

import (
	"strconv"

	"myblog/models"
	"myblog/utils"

	"github.com/beego/beego/v2/client/orm"
)

type ArticleController struct {
	BaseController
}

// @router /article/add [get]
func (this *ArticleController) AddArticleGet() {
	this.TplName = "write_article.html"
}

// @router /article/add [post]
func (this *ArticleController) AddArticlePost() {
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	article := models.Article{
		Title:   title,
		Tags:    tags,
		Short:   short,
		Content: content,
		Author:  "姜峰",
	}
	err := models.ArticleInsert(article)
	if err != nil {
		this.send_json(200, "ok", "True")
	} else {
		this.send_json(500, "插入数据库失败", "False")
	}
}

// @router /article/:id [get]
func (this *ArticleController) ArticleShowGet() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	art := models.QueryArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
	this.TplName = "show_article.html"
}

// @router /article/update [get]
func (this *ArticleController) ArticleUpdateGet() {
	id, _ := this.GetInt("id")
	art := models.QueryArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Tags"] = art.Tags
	this.Data["Short"] = art.Short
	this.Data["Content"] = art.Content
	this.Data["Id"] = art.Id
	this.TplName = "write_article.html"
}

// @router /article/update [post]
func (this *ArticleController) ArticleUpdatePost() {
	id, _ := this.GetInt("id")
	//获取浏览器传输的数据，通过表单的name属性获取值
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	art := models.Article{
		Id:      id,
		Title:   title,
		Tags:    tags,
		Short:   short,
		Content: content,
	}
	o := orm.NewOrm()
	_, err := o.Update(&art)
	if err != nil {
		this.send_json(200, "更新成功", "True")
	} else {
		this.send_json(500, "更新失败", "False")
	}
}

// @router /article/delete [get]
func (this *ArticleController) ArticleDeleteGet() {
	id, _ := this.GetInt("id")
	o := orm.NewOrm()
	o.Delete(&models.Article{Id: id})
	this.Redirect("/", 302)

}

// @router /tags
func (this *ArticleController) TagsGet() {
	this.Data["Tags"] = models.QueryArticleWithParam("tags")
	this.TplName = "tags.html"
}

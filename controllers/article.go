package controllers

import (
	"myblog/models"
	"myblog/utils"
	"strconv"
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

package controllers

import (
	"myblog/models"
)

type HomeController struct {
	BaseController
}

// @router / [get]
func (this *HomeController) HomeGet() {
	tag := this.GetString("tag")
	page, _ := this.GetInt("page")
	var artlist []models.Article
	if len(tag) > 0 {
		artlist = models.QueryArticleWithTag(tag)
		this.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}
		artlist, _ = models.FindArticleWithPage(page)
		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		this.Data["HasFooter"] = true
	}

	this.Data["Content"] = models.MakeHomeBlocks(artlist, this.IsLogin)
	this.TplName = "home.html"
}

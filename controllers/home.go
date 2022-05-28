package controllers

import (
	"fmt"
	"myblog/models"
)

type HomeController struct {
	BaseController
}

// @router / [get]
func (this *HomeController) HomeGet() {
	page, _ := this.GetInt("page")
	if page <= 0 {
		page = 1
	}
	var artList []models.Article
	artList, _ = models.FindArticleWithPage(page)
	this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
	this.Data["HasFooter"] = true
	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)
	this.TplName = "home.html"
}

package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
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
	logs.Error("aaaaaaaaaaaaa")
	this.Data["HasFooter"] = true
	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)
	this.TplName = "home.html"
}

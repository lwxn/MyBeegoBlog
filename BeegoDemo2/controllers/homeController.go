package controllers

import (
	"BeegoDemo2/models"
	"fmt"
)


type HomeController struct {
	BaseController
}

//分页查询
func (this *HomeController)Get(){
	tag := this.GetString("tag")
	var artList []models.Article

	if tag != ""{
		artList = models.GetArticleByTags(tag)
		this.Data["HasFooter"] = false
	}else{
		page, _:= this.GetInt("page")
		if page <= 0{
			page = 1
		}
		fmt.Println(page)
		artList ,_= models.FindArticleWithPage(page)
		this.Data["HasFooter"] = true
		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
	}
	this.Data["Content"] = models.MakeHomeBlocks(artList,this.IsLogin)
	this.TplName = "home.html"
}

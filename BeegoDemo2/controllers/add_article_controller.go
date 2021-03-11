package controllers

import (
	"BeegoDemo2/models"
	"time"
)

type AddArticleController struct{
	BaseController
}


func (this *AddArticleController)Get(){
	this.TplName = "write_article.html"
}

func (this *AddArticleController) Post(){
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	_,err := models.AddArticle(models.Article{0,title,tags,short,content,
		"lwxn",time.Now().Unix()})
	if err != nil{
		this.Data["json"] = map[string]interface{}{"code":"0","message":"插入文章错误"}
	}else{
		this.Data["json"] = map[string]interface{}{"code":"1","message":"插入成功"}
	}
	this.ServeJSON()
}

package controllers

import (
	"BeegoDemo2/models"
	"BeegoDemo2/utils"
	"github.com/astaxie/beego"
	"strconv"
)

type ShowArticleController struct {
	beego.Controller
}

func (this *ShowArticleController)Get(){
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	//按照ID查文章
	art := models.FindArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
	//显示页面
	this.TplName = "show_article.html"
}

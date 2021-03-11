package controllers

import (
	"BeegoDemo2/models"
	"github.com/astaxie/beego"
)

type TagController struct {
	beego.Controller
}

func (this *TagController)Get(){
	this.Data["Tags"] = models.GetMapOfTag()
	this.TplName = "tags.html"
}

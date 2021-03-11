package controllers

import "github.com/astaxie/beego"

type AboutMeController struct {
	beego.Controller
}

func (this *AboutMeController)Get(){
	this.Data["wechat"] = "zhouzishuwolaopo"
	this.Data["qq"] = "zhouzishuwolaopo"
	this.Data["tel"] = "520"

	this.TplName = "aboutme.html"
}

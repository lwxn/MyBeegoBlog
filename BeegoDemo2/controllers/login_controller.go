package controllers

import (
	"BeegoDemo2/models"
	"BeegoDemo2/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct{
	beego.Controller
}

func (this *LoginController) Get(){
	this.TplName = "login.html"
}

//post:login是为了确认真的有这个用户
func (this *LoginController) Post(){
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println(username,password)

	password = utils.MD5(password)
	id := models.QueryUserWithPassword(username,password)
	if id > 0{
		//设置session,之后就可以用cookie来判断用户是谁了
		this.SetSession("loginuser",username)
		this.Data["json"] = map[string]interface{}{"code":"1","message":"登录成功"}
	}else{
		this.Data["json"] = map[string]interface{}{"code":"0","message":"登录失败"}
	}
	this.ServeJSON()
}

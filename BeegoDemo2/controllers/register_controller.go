package controllers

import (
	"BeegoDemo2/models"
	"BeegoDemo2/utils"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type RegisterController struct{
	beego.Controller
}

func (this *RegisterController) Get(){
	this.TplName = "register.html"
}

//处理注册
func (this *RegisterController) Post(){
	//获取表单信息
	username := this.GetString("username")
	password := this.GetString("password")

	//看一下user是否已经存在了
	id := models.QueryUserWithUsername(username)
	fmt.Println(id)
	if id > 0{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"用户名已存在"}
		this.ServeJSON()
		return
	}

	password = utils.MD5(password)
	user := models.User{0,username,password,0,time.Now().Unix()}
	_,err := models.InsertUser(user)
	if err != nil{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"the user has been existed!!"}
	}else{
		this.Data["json"] = map[string]interface{}{"code":1,"message":"User " + username + " register success"}
	}

	this.ServeJSON()
}


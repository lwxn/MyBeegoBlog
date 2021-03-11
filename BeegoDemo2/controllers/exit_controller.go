package controllers

type ExitController struct {
	BaseController
}


func (this *ExitController)Get(){
	//清除用户数据
	this.DelSession("loginuser")
	this.Redirect("/",302)
}
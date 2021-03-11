package controllers

import (
	"BeegoDemo2/models"
	"fmt"
	"github.com/astaxie/beego"
)

type DeleteArticleController struct {
	beego.Controller
}

func (this *DeleteArticleController)Get(){
	id,_ := this.GetInt("id")
	_, err := models.DeleteArticleWithId(id)

	if err != nil{
		fmt.Println(err)
	}
	this.Redirect("/",302)
}
package controllers

import (
	"BeegoDemo2/models"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type UpdateArticlecontroller struct {
	beego.Controller
}

func (this *UpdateArticlecontroller)Get(){
	id,_ := this.GetInt("id")
	fmt.Println("id",id)
	article := models.FindArticleWithId(id)

	//修改html数据
	this.Data["Title"] = article.Title
	this.Data["Tags"] = article.Tags
	this.Data["Short"] = article.Short
	this.Data["Content"] = article.Content
	this.Data["Id"] = article.Id
	//返回页面
	this.TplName = "write_article.html"
}

//Post
func (this *UpdateArticlecontroller)Post()  {
	id,_ := this.GetInt("id")
	fmt.Println("Post:","id",id)

	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	fmt.Println("Post","title",title)
	fmt.Println("Post","tags",tags)
	fmt.Println("Post","short",short)
	fmt.Println("Post","content",content)


	new_article := models.Article{id,title,tags,short,content,"",time.Now().Unix()}
	_,err := models.UpdateArticleWithNew(new_article)
	if err != nil{
		this.Data["json"] = map[string]interface{}{"code":"0","message":"Update fail!"}
	}else{
		this.Data["json"] = map[string]interface{}{"code":"1","message":"Update success!"}
	}
	this.ServeJSON()
}

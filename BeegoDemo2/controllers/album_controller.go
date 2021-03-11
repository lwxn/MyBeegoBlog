package controllers

import (
	"BeegoDemo2/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AlbumController struct {
	beego.Controller
}

func (this *AlbumController)Get(){
	albums,err := models.FindAllAlbums()
	if err != nil{
		logs.Error(err)
		return
	}
	this.Data["Album"] = albums
	this.TplName = "album.html"
}


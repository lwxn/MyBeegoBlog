package controllers

import (
	"BeegoDemo2/models"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type UploadController struct {
	beego.Controller
}


func (this *UploadController)Post()  {
	fmt.Println("fileupload...")
	fileData,fileHeader,err := this.GetFile("upload")
	if err != nil{
		this.responseErr(err)
		return
	}
	fmt.Println("name:",fileHeader.Filename,fileHeader.Size)
	fmt.Println(fileData)

	//通过后缀判断是否图片
	fileType := "other"
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg"{
		fileType = "img"
	}else{
		log.Println("这个不是图片啊")
		return
	}

	//图片的本地存储路径
	fileDir := fmt.Sprintf("static/img/%s/%d/%d/%d",fileType,time.Now().Year(),
		time.Now().Month(),time.Now().Day())
	err = os.MkdirAll(fileDir,os.ModePerm)
	if err != nil{
		this.responseErr(err)
		return
	}
	//copy image
	fmt.Println("copy image begin!")
	fileName := fmt.Sprintf(fileHeader.Filename)
	filePathStr := filepath.Join(fileDir,fileName)
	desFile,err:= os.Create(filePathStr)
	if err != nil{
		this.responseErr(err)
		return
	}
	_,err = io.Copy(desFile,fileData)
	if err != nil{
		this.responseErr(err)
		return
	}
	models.InsertAlbum(models.Album{0,filePathStr,
		fileName,0,time.Now().Unix()})
	this.Data["json"] = map[string]interface{}{"code":"1","message":"upload the img successful!"}
	this.ServeJSON()
}

func (this *UploadController)responseErr(err error)(){
	this.Data["json"] = map[string]interface{}{"message:":err}
	this.ServeJSON()
}

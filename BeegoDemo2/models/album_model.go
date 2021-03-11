package models

import (
	"BeegoDemo2/utils"
	"fmt"
)

type Album struct {
	Id int
	Filepath string
	Filename string
	Status int
	Createtime int64
}


func InsertAlbum(album Album){
	_,err := utils.ModifyDB("insert into album(filepath,filename,status,createtime)" +
		"values(?,?,?,?)",album.Filepath,album.Filename,album.Status,album.Createtime)
	//_,err := utils.ModifyDB(sql)
	if err != nil{
		fmt.Println(err)
	}
}

func FindAllAlbums() ([]Album, error) {
	sql := "select id,filepath,filename,status,createtime from album;"
	rows,err := utils.QueryDB(sql)
	if err != nil{
		fmt.Println("查询图片失败！")
		return nil,err
	}

	var albums []Album
	for rows.Next(){
		var(
			id int
			filepath string
			filename string
			status int
			createtime int64
		)
		rows.Scan(&id,&filepath,&filename,&status,&createtime)
		albums = append(albums,Album{id,filepath,filename,status,createtime})
	}
	return albums,nil
}

package models

import (
	"BeegoDemo2/utils"
	"fmt"
)

type User struct {
	Id int
	UserName string
	Password string
	Status int
	CreateTime int64
}

//type User struct{
//	Id int
//	UserName string
//	Password string
//	Status int
//	CreateTime int64
//}

//数据库操作

//插入用户
func InsertUser(user User)(int64,error){
	return utils.ModifyDB("insert into users(username,password,status,createTime) values (?,?,?,?)",
		user.UserName,user.Password,user.Status,user.CreateTime)
}

//带有条件的查询语句
func QueryUserWithCon(con string)int{
	sql := fmt.Sprintf("select id from users %s",con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名来查询
func QueryUserWithUsername(username string)int{
	sql := fmt.Sprintf("where username = '%s'",username)
	return QueryUserWithCon(sql)
}

func QueryUserWithPassword(username string, password string)int{
	sql := fmt.Sprintf("where username = '%s' and password = '%s'",username,password)
	return QueryUserWithCon(sql)
}



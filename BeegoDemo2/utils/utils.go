package utils

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB
// insert into user (?,?,?) values(1,2,3) 并且返回改了多少条
func ModifyDB(sql string, args ...interface{})(int64,error){
	result,err :=db.Exec(sql,args...)
	if err != nil{
		log.Println(err)
		return 0,err
	}

	count, err := result.RowsAffected()
	if err !=nil{
		log.Println(err)
		return 0,err
	}
	return count,err

}


func QueryRowDB(sql string) *sql.Row{
	return db.QueryRow(sql)
}

func QueryDB(sql string)(*sql.Rows,error){
	return db.Query(sql)
}

//table:users  id,username,password,status,createTime
func CreateTableWithUser(){
	sql := `CREATE TABLE IF NOT EXISTS users(
			id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
			username VARCHAR(64),
			password VARCHAR(64),
			status INT(4),
			createTime INT(10)
		);`
	ModifyDB(sql)
}



func InitMysql() {
	fmt.Println("Init Mysql")
	driverName := beego.AppConfig.String("driverName")

	//注册数据库driver
	//orm.RegisterDriver(driverName,orm.DRMySQL)

	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	//root:root@tcp(127.0.0.1:3306)/myblog?charset=utf8
	dbConn := user + ":" + pwd + "@tcp(" + host +":" + port + ")/" + dbname +
		"?charset=utf8"

	db1, err := sql.Open(driverName,dbConn)
	fmt.Println("open connection successful!")
	if err != nil{
		fmt.Println(err.Error())
	} else{
		db = db1
		CreateTableWithUser()
		CreateTableWithArticle()
		CreateTableWithAlbum()
	}

}


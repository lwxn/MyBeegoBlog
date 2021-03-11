package main

import (
	_ "BeegoDemo2/routers"
	"BeegoDemo2/utils"
	"github.com/astaxie/beego"
)

func main() {
	utils.InitMysql()
	beego.Run()
}
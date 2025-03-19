package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "myblog/routers"
)

func main() {

	//utils.InitMysql()
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	dbconn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dbconn)
	beego.Run()
}

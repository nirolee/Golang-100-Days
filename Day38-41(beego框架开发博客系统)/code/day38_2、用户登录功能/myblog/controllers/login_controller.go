package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"myblog/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:", username, ",password:", password)
	hash := md5.Sum([]byte(password))
	password = hex.EncodeToString(hash[:])
	id := models.QueryUserWithParam(username, password)
	fmt.Println("id:", id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	this.ServeJSON()
}

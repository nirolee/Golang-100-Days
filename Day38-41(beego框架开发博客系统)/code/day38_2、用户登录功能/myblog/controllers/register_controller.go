package controllers

import (
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"myblog/models"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

func (this *RegisterController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	if username == "" || password == "" {
		this.Data["json"] = map[string]string{"error": "Missing username or password"}
		this.Ctx.Output.SetStatus(400)
		this.ServeJSON()
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		this.Data["json"] = map[string]string{"error": "Failed to encrypt password"}
		this.Ctx.Output.SetStatus(500)
		this.ServeJSON()
		return
	}
	existUser, err := models.GetUsersByName(username)
	if err == nil && existUser != nil {
		this.Data["json"] = "用户名已存在"
		this.ServeJSON()
		return
	}
	user := models.Users{
		Username:   username,
		Password:   string(hashedPassword),
		Status:     1,
		Createtime: int(time.Now().Unix()),
	}
	if _, err := models.AddUsers(&user); err != nil {
		this.Data["json"] = map[string]string{"error": "Failed to insert user"}
		this.Ctx.Output.SetStatus(500)
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]string{"success": "Register successfully"}
	this.ServeJSON()
}

// 处理注册
//func (this *RegisterController) Post1() {
//	log := logrus.New()
//	//获取表单信息
//	username := this.GetString("username")
//	password := this.GetString("password")
//	repassword := this.GetString("repassword")
//	fmt.Println(username, password, repassword)
//	log.Info(username, password, repassword)
//
//	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
//	id := models.QueryUserWithUsername(username)
//	fmt.Println("id:", id)
//	if id > 0 {
//		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已经存在"}
//		this.ServeJSON()
//		return
//	}
//
//	//注册用户名和密码
//	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
//
//	hash := md5.Sum([]byte(password))
//	password = hex.EncodeToString(hash[:])
//
//	fmt.Println("md5后：", password)
//
//	user := models.User{0, username, password, 0, time.Now().Unix()}
//	_, err := models.InsertUser(user)
//	if err != nil {
//		this.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
//	} else {
//		this.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
//	}
//	this.ServeJSON()
//
//}

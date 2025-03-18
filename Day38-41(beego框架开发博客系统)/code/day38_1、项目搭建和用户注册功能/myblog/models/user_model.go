package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"myblog/utils"
	"time"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态， 1删除
	Createtime int64
}

func init() {
	orm.RegisterModel(new(User))
}

//--------------数据库操作-----------------

// 插入
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

// 按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

// 根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

// 根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}

func AddUser() {
	o := orm.NewOrm()
	user := User{Username: "admin", Password: "admin", Status: 0, Createtime: time.Now().Unix()}
	_, err := o.Insert(&user)
	if err != nil {
		fmt.Println("插入失败", err)
	} else {
		fmt.Println("插入成功")
	}
}

func GetUserById(id int) User {
	o := orm.NewOrm()
	user := User{Id: id}
	o.Read(&user, "Id")
	return user
}

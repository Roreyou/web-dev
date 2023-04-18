package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //前面加下划线表示没有直接使用
)

// 定义全局变量
// 定义打开数据库并连接等操作
// 定义模型，增删改查操作
// 函数名首字母都必须大写
type User struct { //定义数据库模型
	ID       int64
	Username string `json:"username"` //用户名
	Password string `json:"password"`
}

func Openmysql() (db1 *gorm.DB) { //打开数据库
	//连接MYSQL数据库
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1)/testsystem?charset-utf8mb4&parseTime=True&loc=Local")
	//parseTime表示将数据库中时间类型的解析为时间，loc表示解析本地时间
	if err != nil { //判断是否打开错误
		panic(err)
	}
	//2、把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})
	//创建一个用户便于测试
	// u1 := User{1, "小王子", "123"}
	// db.Create(&u1)
	return db
}

func FindUsername(name string) (user User) { //查找到用户名相同的用户
	var auser User
	//查找输入用户名是否存在
	db := Openmysql()
	db.Where("username = ?", name).First(&auser)
	fmt.Printf("u:%#v\n", auser)
	return auser
}

func CreateaUser(name string, password string) (auser User) { //创建一个用户
	db := Openmysql() //获取数据库对象
	user := User{Username: name, Password: password}
	db.Create(&user)
	return user
}

package dao

import (
	_ "github.com/jinzhu/gorm/dialects/mysql" //前面加下划线表示没有直接使用
)

// 定义全局变量
// 定义打开数据库并连接等操作
// 定义模型，增删改查操作
// 函数名首字母都必须大写
/* 不适用 使用FindUser算了
func FindUserId(name string) (user User_Info) { //查找到用户名相同的用户
	var auser User_Info
	//查找输入用户名是否存在
	db := Openmysql()
	//2、把模型与数据库中的表对应起来
	db.AutoMigrate(&User_Info{})
	db.Where("user_id = ?", name).First(&auser)
	fmt.Printf("u:%#v\n", auser)
	return auser
}
*/
func CreateaUser(name string, password string) (auser User_Info) { //创建一个用户
	db := Openmysql() //获取数据库对象
	//2、把模型与数据库中的表对应起来
	db.AutoMigrate(&User_Info{})
	user := User_Info{User_id: 100, User_Name: "测试姓名", Real_name: "测试真实姓名", User_password: "测试密码"}
	db.Create(&user)
	return user
}

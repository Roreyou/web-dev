package dao

import "github.com/jinzhu/gorm"

func Openmysql() (db1 *gorm.DB) { //打开数据库
	//连接MYSQL数据库
	db, err := gorm.Open("mysql", "root:0921@(127.0.0.1)/web_database?charset-utf8mb4&parseTime=True&loc=Local")
	//parseTime表示将数据库中时间类型的解析为时间，loc表示解析本地时间
	if err != nil { //判断是否打开错误
		panic(err)
	}
	//defer db.Close() //把数据库的连接关闭掉
	//2、把模型与数据库中的表对应起来
	//db.AutoMigrate(&User{})
	//创建一个用户便于测试
	//u1 := User{1, "小王子", "123"}
	//db.Create(&u1)
	return db
}

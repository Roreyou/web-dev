package dao

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AdminInfo struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	ActRecord ActRecord `json:"actrecord"` //操作记录的结构体
}
type ActRecord struct {
	Test string //还没想清楚记录什么
}

func GetAdminInfo(c *gin.Context) *AdminInfo {
	db, err := gorm.Open("mysql", "root:0921@(127.0.0.1:3306)/testsystem?charset=utf8mb4&parseTime=True&loc=Local")
	// db就是database
	if err != nil {
		panic(err)
	}
	defer db.Close() //把数据库的连接关闭掉

	db.AutoMigrate(&AdminInfo{})
	u1_record := ActRecord{"在一月一日租用了a号服务器，当前使用时长xxx"}
	u2_record := ActRecord{"在一月二日租用了b号服务器，当前使用时长xxx"}
	u1 := AdminInfo{1, "小明", "1111", u1_record}
	u2 := AdminInfo{2, "小红", "1111", u2_record}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	fmt.Println("完成！")

	return &u1 //用来测试一下 返回的是u1数据
}

package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Gpu struct {
	ID   int    `json:"机器编号"`
	Type string `json:"型号"`
	Size int    `json:"显存大小"`
	Flag bool   `json:"是否被租用"`
	User string `json:"租用用户"`
} //操作gpu信息的结构体

func GetGpuInfo(c *gin.Context) *Gpu {
	//连接sql数据库
	db, err := gorm.Open("mysql", "root:0921@(127.0.0.1:3306)/gpudata?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //把数据库的连接关闭掉

	db.AutoMigrate(&Gpu{}) //自动迁移，使结构体何数据库对应
	//u1 := Gpu{ID: 1, Type: "RTX3080", Size: 10, Flag: false, User: ""}
	//db.Create(&u1)//增加数据
	var u2 Gpu
	db.First(&u2) //查询数据
	//db.Model(&u2).Update("flag", false) //更新数据
	return &u2 //返回数据
}

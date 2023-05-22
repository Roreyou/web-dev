package dao

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Gpu struct {
	Server_id    int    `json:"server_id"`
	Server_type  string `json:"sever_type"`
	Server_size  int    `json:"server_size"`
	Server_state int    `json:"server_state"`
	Server_flag  bool   `json:"sever_flag"`
} //操作gpu信息的结构体

func GetGpuInfo(c *gin.Context) *Gpu {
	//连接sql数据库
	db := Openmysql()
	db.AutoMigrate(&Gpu{}) //自动迁移，使结构体何数据库对应
	//u1 := Gpu{Server_id: 1, Server_type: "RTX3080", Server_size: 10, Server_state: 0, Server_flag: false}
	//db.Create(&u1) //增加数据
	var u2 Gpu
	db.First(&u2) //查询数据
	//db.Model(&u2).Update("flag", false) //更新数据
	return &u2 //返回数据
}

package dao

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server_Info struct {
	Server_id    int    `json:"server_id"`
	Server_type  string `json:"sever_type"`
	Server_size  int    `json:"server_size"`
	Server_state int    `json:"server_state"`
	Server_flag  bool   `json:"sever_flag"`
} //操作gpu信息的结构体

func (Server_Info) TableName() string {
	return "server_info"
}

func GetGpuInfo(c *gin.Context) *[]Server_Info {
	//连接sql数据库
	db := Openmysql()
	db.AutoMigrate(&Server_Info{}) //自动迁移，使结构体何数据库对应
	var u2 []Server_Info
	db.Find(&u2) //查询数据
	return &u2   //返回数据
}

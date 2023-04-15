package dao

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Admin struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	RoleId    int64  `json:"role_id"`
	ActRecord string `json:"actrecord"`
}

func adminInfo(c *gin.Context) *Admin {
	admin := &Admin{} // 用来创建相关信息
	return admin
}

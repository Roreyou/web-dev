package controller

import (
	"net/http"

	"back/dao"

	"github.com/gin-gonic/gin"
)

func InfoAdminControl(c *gin.Context) {
	//该函数用于调用dao里面封装的用于调用数据信息的函数
	c.JSON(http.StatusOK, dao.GetAdminInfo(c))
}

package controller

import (
	"net/http"

	"back/dao"

	"github.com/gin-gonic/gin"
)

func InfoUserControl(c *gin.Context) {
	/*
		//该函数用于调用dao里面封装的用于调用数据信息的函数
		_, err := middlewares.GetToken(c) //获取token
		if err != nil {
			//c.AbortWithError(500, err)
			c.HTML(http.StatusForbidden, "no_right.html", nil)
		} else {
			c.JSON(http.StatusOK, dao.GetAdminInfo(c))
		}
	*/
	//不用token的测试版本
	c.JSON(http.StatusOK, dao.GetUserInfo(c))
}

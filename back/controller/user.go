package controller

import (
	"net/http"

	"back/dao"

	"github.com/gin-gonic/gin"
)

func InfoUserControl(c *gin.Context) {
	/*
		//用token的测试版本
		_, err := middlewares.GetToken(c) //获取token
		if err != nil {
			//c.AbortWithError(500, err)
			c.HTML(http.StatusForbidden, "no_right.html", nil)
		} else {
			c.JSON(http.StatusOK, dao.GetUserInfo(c))
		}
	*/
	//不用token的测试版本
	c.JSON(http.StatusOK, dao.GetUserInfo(c))
}

func PasswordUserControl(c *gin.Context) {
	/*
		//用token的测试版本
		_, err := middlewares.GetToken(c) //获取token
		if err != nil {
			//c.AbortWithError(500, err)
			c.HTML(http.StatusForbidden, "no_right.html", nil)
		} else {
			if_success := dao.ChangeUserPassword(c)
			if if_success {
				c.JSON(http.StatusOK, gin.H{
					"Success": "修改成功！",
				})
			} else {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"fail": "原密码错误",
				})
			}
		}
	*/
	//不用token的测试版本
	if_success := dao.ChangeUserPassword(c)

	if if_success {
		c.JSON(http.StatusOK, gin.H{
			"Success": "修改成功！",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"fail": "原密码错误",
		})
	}
}

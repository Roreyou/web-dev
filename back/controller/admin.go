package controller

import (
	"back/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitPassword_Controller(c *gin.Context) {
	/*
		//用token的测试版本
		_, err := middlewares.GetToken(c) //获取token
		if err != nil {
			c.HTML(http.StatusForbidden, "no_right.html", nil)
		} else {
			if_success := dao.Init_password(c)
			if if_success == true {
				c.JSON(http.StatusOK, gin.H{
					"Success": "修改成功！",
				})
			} else {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"fail": "用户不存在！",
				})
			}
		}
	*/
	//不用token的测试版本
	if_success := dao.Init_password(c)
	if if_success == true {
		c.JSON(http.StatusOK, gin.H{
			"Success": "修改成功！",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"fail": "用户不存在！",
		})
	}
}

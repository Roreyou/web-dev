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

	get_info := dao.GetUserInfo(c)
	if get_info.User_id == 0 { //该用户不存在则报错
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"fail": "该用户名不存在",
		})
	} else {
		c.JSON(http.StatusOK, get_info)
	}

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

func Record_controller(c *gin.Context) {
	/*
		//用token的测试版本
		_, err := middlewares.GetToken(c) //获取token
		if err != nil {
			//c.AbortWithError(500, err)
			c.HTML(http.StatusForbidden, "no_right.html", nil)
		} else {
			c.JSON(http.StatusOK, dao.Get_recording(c))
		}
	*/
	//不用token的测试版本

	recording_info := dao.Get_Recording(c)
	c.JSON(http.StatusOK, recording_info)

}

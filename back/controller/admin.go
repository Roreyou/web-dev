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
	if if_success {
		c.JSON(http.StatusOK, gin.H{
			"Success": "修改成功！",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"fail": "用户不存在！",
		})
	}
}

func Delete_usrContrller(c *gin.Context) {
	/*
		_, err := middlewares.GetToken(c) //获取token
		if err != nil {
			c.HTML(http.StatusForbidden, "no_right.html", nil)
		} else {
			if_success := dao.Delete_usr(c)
			if if_success == true{
				c.JSON(http.StatusOK, gin.H{
					"success": "删除成功！",
				})
			}else{
				c.JSON(http.StatusOK, gin.H{
					"fail": "用户不存在！",
				})
			}
		}*/
	if_success := dao.Delete_usr(c)
	if if_success {
		c.JSON(http.StatusOK, gin.H{
			"Success": "删除成功！",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"fail": "用户不存在！",
		})
	}

}

func Add_usrContrller(c *gin.Context) {
	/*
		_, err := middlewares.GetToken(c)
		if err != nil {
			c.HTML(http.StatusForbidden, "no_right.html", nil)
		} else {
			if_success := dao.Add_user(c)
			if if_success == true {
				c.JSON(http.StatusOK, gin.H{
					"success": "增加成功！",
				})
			}
		}*/
	if_success := dao.Add_user(c)
	if if_success {
		c.JSON(http.StatusOK, gin.H{
			"Success": "增加成功！",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"fail": "用户不存在！",
		})
	}

}

func Show_alluserContrller(c *gin.Context) {
	/*
		_, err := middlewares.GetToken(c)
			if err != nil {
				c.HTML(http.StatusForbidden, "no_right.html", nil)
			} else {
				c.JSON(http.StatusOK, dao.Show_alluser(c))
			}
	*/
	c.JSON(http.StatusOK, dao.Show_alluser(c))
}

func FindAllContController(c *gin.Context) {
	/*
		//用token的测试版本
		_, err := middlewares.GetToken(c) //获取token
		if err != nil {
			//c.AbortWithError(500, err)
			c.HTML(http.StatusForbidden, "no_right.html", nil)
		} else {
			c.JSON(http.StatusOK, dao.FindAllCont())
		}*/
	c.JSON(http.StatusOK, dao.FindAllCont())
}

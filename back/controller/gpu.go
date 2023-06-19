package controller

import (
	"net/http"

	"back/dao"

	"back/services"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	ssg := dao.GetGpuInfo(c)
	if ssg == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "当前无可用机器",
		})

	} else {
		c.JSON(http.StatusOK, ssg)
	}
}

func Link(c *gin.Context) {
	user_id := c.PostForm("user_id")
	machineid := c.PostForm("server_id")
	image_id := c.PostForm("image_id")
	user_password := c.PostForm("user_password")
	result := services.GreateDocker(machineid, user_id, image_id, user_password)
	if result == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Success": "容器创建成功",
		})
	}
	if result == 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "时间已超额",
		})
	}
	if result == 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"fail": "容器创建失败",
		})
	}
}

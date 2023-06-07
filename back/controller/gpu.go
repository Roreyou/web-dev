package controller

import (
	"fmt"
	"net/http"

	"back/dao"

	"back/services"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	c.JSON(http.StatusOK, dao.GetGpuInfo(c))
}

func Link(c *gin.Context) {
	machineid := c.Query("server_id")
	user_id := c.Query("user_id")
	docker_id := c.Query("docker_id")
	fmt.Println(machineid)
	fmt.Println(user_id)
	fmt.Println(docker_id)
	result := services.GreateDocker(machineid, user_id, docker_id)
	if result == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Success": "创建成功",
		})
	}
	if result == 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "时间已超额",
		})
	}
}

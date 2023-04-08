package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("api") // 创建一个名叫api的路由组,并且该组的统一前缀为/api
	api.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "您已经进入了登陆页面，在这个页面，你可以输入你的个人信息，以便进行_进入后面页面前的_校验"}) //不需要权限校验的页面
	})

	router.Run(":8080")
}

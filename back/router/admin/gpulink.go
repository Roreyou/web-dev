package admin

import (
	"back/controller"
	"back/middlewares"

	"github.com/gin-gonic/gin"
)

func GpulinkRoute(r *gin.Engine) {
	linkrouter := r.Group("/linkgpu")
	linkrouter.Use(middlewares.AuthToken()) //须认证
	{
		linkrouter.GET("/item", controller.Show)  //获取可用机器列表
		linkrouter.POST("/link", controller.Link) //将用户与对应机器相连
	}

}

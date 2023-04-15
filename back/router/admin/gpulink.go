package admin

import (
	AdminControllers "back/controller"
	"back/middlewares"

	"github.com/gin-gonic/gin"
)

func GpulinkRoute(r *gin.Engine) {
	linkrouter := r.Group("/linkgpu")
	gpuC := &AdminControllers.Gpu{}         //控制器
	linkrouter.Use(middlewares.AuthToken()) //须认证
	{
		linkrouter.GET("/item", gpuC.Show)  //获取可用机器列表
		linkrouter.POST("/link", gpuC.Link) //将用户与对应机器相连
	}

}

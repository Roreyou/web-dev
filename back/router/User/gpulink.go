package User

import (
	"back/controller"
	"back/middlewares"

	"github.com/gin-gonic/gin"
)

func GpulinkRoute(r *gin.Engine) {
	linkrouter := r.Group("/user")
	linkrouter.Use(middlewares.AuthToken()) //须认证
	{
		linkrouter.GET("/get_item", controller.Show)    //获取可用机器列表
		linkrouter.POST("/create_gpu", controller.Link) //创建服务器
	}

}

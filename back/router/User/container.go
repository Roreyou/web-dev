package User

import (
	"back/controller"

	"github.com/gin-gonic/gin"
)

func EnterContainerRoute(r *gin.Engine) { //开机
	//UserRouter := r.Group("/user")
	r.POST("/work", controller.ConnectContainer)
}

func ExitContainerRoute(r *gin.Engine) { //删除服务器
	//UserRouter := r.Group("/user")
	r.POST("/exit", controller.ExitContainer)
}

func DeleteContainerRoute(r *gin.Engine) { //关机
	r.GET("/close", controller.DeleteContainer)
}

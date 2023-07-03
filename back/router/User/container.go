package User

import (
	"back/controller"

	"github.com/gin-gonic/gin"
)

func EnterContainerRoute(r *gin.Engine) { //开机
	UserRouter := r.Group("/user")
	UserRouter.POST("/work", controller.ConnectContainer)
}

func ExitContainerRoute(r *gin.Engine) { //删除服务器
	UserRouter := r.Group("/user")
	UserRouter.POST("/exit", controller.ExitContainer)
}

func DeleteContainerRoute(r *gin.Engine) { //关机
	UserRouter := r.Group("/user")
	UserRouter.POST("/close", controller.DeleteContainer)
}

func ContainerInfo(r *gin.Engine) { //展示容器信息
	UserRouter := r.Group("/user")
	UserRouter.POST("/show", controller.ShowContainerInfo)
}

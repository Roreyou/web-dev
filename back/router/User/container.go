package User

import (
	"back/controller"

	"github.com/gin-gonic/gin"
)

func EnterContainerRoute(r *gin.Engine) {
	UserRouter := r.Group("/user")
	UserRouter.POST("/work", controller.ConnectContainer)
}

func ExitContainerRoute(r *gin.Engine) {
	UserRouter := r.Group("/user")
	UserRouter.POST("/exit", controller.ExitContainer)
}

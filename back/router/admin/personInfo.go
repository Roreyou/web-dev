package admin

import (
	"back/controller"
	"back/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminInfoRoute(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	adminRouter.Use(middlewares.AuthToken())              //调用中间件，用于进行权限认证
	adminRouter.GET("/info", controller.InfoAdminControl) //开始响应，在我们的控制器文件中调用我的处理函数
}

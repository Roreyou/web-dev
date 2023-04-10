package admin

import (
	AdminControllers "back/controller"
	"back/middlewares"

	"github.com/gin-gonic/gin"
)

func PersonInfoRoute(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	adminRouter.Use(middlewares.AuthToken()) //调用中间件，用于进行权限认证
	{
		adminRouter.GET("/info", AdminControllers.InfoAdmin) //使用控制器 用于调用admin的信息
	}
}

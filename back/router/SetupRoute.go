package router

import (
	"back/middlewares"
	"back/router/User"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // 创建一个gin引擎
	r.Use(middlewares.AuthToken())
	GethtmlRouter(r)
	CheckUserRouter(r)
	User.UserInfoRoute(r)
	User.User_changePasswordRoute(r)
	User.GpulinkRoute(r)
	User.EnterContainerRoute(r)
	User.DeleteContainerRoute(r)
	User.ExitContainerRoute(r)
	return r
}

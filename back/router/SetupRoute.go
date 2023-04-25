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
	User.UserInfoRoute(r) // 调用我的adminIndoRoute来注册中间件auth，并创建一个路由，访问admin的个人信息
	User.GpulinkRoute(r)
	return r
}

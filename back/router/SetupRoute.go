package router

import (
	"back/router/admin"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()      // 创建一个gin引擎
	admin.AdminInfoRoute(r) // 调用我的adminIndoRoute来注册中间件auth，并创建一个路由，访问admin的个人信息

	return r
}

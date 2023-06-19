package router

import (
	"back/middlewares"
	"back/router/Admin"
	"back/router/User"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // 创建一个gin引擎
	r.Use(middlewares.AuthToken())
	//r.Use(middlewares.CorsHander())
	GethtmlRouter(r)
	CheckUserRouter(r)
	User.UserInfoRoute(r)
	User.User_changePasswordRoute(r)
	User.User_recordingRoute(r)
	User.GpulinkRoute(r)
	Admin.InitPasswordRoute(r)
	Admin.Add_usrRoute(r)
	Admin.Delete_usrRoute(r)
	Admin.Show_usrRoute(r)
	Admin.Show_contRoute(r)
	User.EnterContainerRoute(r)
	User.DeleteContainerRoute(r)
	User.ExitContainerRoute(r)
	User.ContainerInfo(r)
	return r
}

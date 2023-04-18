package router

import (
	"back/controller"

	"github.com/gin-gonic/gin"
)

func GethtmlRouter(r *gin.Engine) { //显示登陆页面
	r.LoadHTMLFiles("./login.html") //将html文件加载进入
	r.GET("/login", controller.Showhtml)
}

func CheckUserRouter(r *gin.Engine) { //登陆验证
	r.POST("/enter", controller.LoginInfoController)
}

func TestMiddlerRouter(r *gin.Engine) { //测试中间件
	r.POST("/doing", controller.TestMiddler)
}

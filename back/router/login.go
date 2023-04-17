package router

import (
	"back/controller"

	"github.com/gin-gonic/gin"
)

func Gethtml(r *gin.Engine) { //显示登陆页面
	r.GET("/login", controller.Showhtml)
}
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html") //将html文件加载进入
	return r
}
func CheckUser(r *gin.Engine) { //登陆验证
	r.POST("/enter", controller.LoginInfoController)
}

func CreateUser(r *gin.Engine) { //注册
	r.Post("/logon", controller.logon)
}

package User

import (
	"back/controller"
	// "back/middlewares"

	"github.com/gin-gonic/gin"
)

func UserInfoRoute(r *gin.Engine) {
	UserRouter := r.Group("/user")
	r.LoadHTMLFiles("./no_right.html")                   //调用中间件，用于进行权限认证
	UserRouter.POST("/info", controller.InfoUserControl) //开始响应，在我们的控制器文件中调用我的处理函数
}

func User_changePasswordRoute(r *gin.Engine) {
	UserRouter := r.Group("/user")
	r.LoadHTMLFiles("./no_right.html")                                  //调用中间件，用于进行权限认证
	UserRouter.POST("/change_password", controller.PasswordUserControl) //开始响应，在我们的控制器文件中调用我的处理函数
}

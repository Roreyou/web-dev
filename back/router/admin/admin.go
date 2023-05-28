package Admin

import (
	"back/controller"

	"github.com/gin-gonic/gin"
)

func InitPasswordRoute(r *gin.Engine) {
	InitPRouter := r.Group("/admin")
	r.LoadHTMLFiles("./no_right.html")                                     //调用中间件，用于进行权限认证
	InitPRouter.POST("/init_password", controller.InitPassword_Controller) //开始响应，在我们的控制器文件中调用我的处理函数
}

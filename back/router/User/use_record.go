package User

import (
	"back/controller"

	"github.com/gin-gonic/gin"
)

func User_recordingRoute(r *gin.Engine) {
	UserRouter := r.Group("/user")
	r.LoadHTMLFiles("./no_right.html")                          //调用中间件，用于进行权限认证
	UserRouter.POST("/recording", controller.Record_controller) //开始响应，在我们的控制器文件中调用我的处理函数
}

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

func Delete_usrRoute(r *gin.Engine) {
	InitPRouter := r.Group("/admin")
	r.LoadHTMLFiles("./no_right.html")                               //调用中间件，用于进行权限认证
	InitPRouter.POST("/delete_user", controller.Delete_usrContrller) //开始响应，在我们的控制器文件中调用我的处理函数
}

func Add_usrRoute(r *gin.Engine) {
	InitPRouter := r.Group("/admin")
	r.LoadHTMLFiles("./no_right.html")                         //调用中间件，用于进行权限认证
	InitPRouter.POST("/add_user", controller.Add_usrContrller) //开始响应，在我们的控制器文件中调用我的处理函数
}

func Show_usrRoute(r *gin.Engine) {
	InitPRouter := r.Group("/admin")
	r.LoadHTMLFiles("./no_right.html")                               //调用中间件，用于进行权限认证
	InitPRouter.POST("/show_user", controller.Show_alluserContrller) //开始响应，在我们的控制器文件中调用我的处理函数
}

func Show_contRoute(r *gin.Engine) {
	InitPRouter := r.Group("/admin")
	r.LoadHTMLFiles("./no_right.html")                                   //调用中间件，用于进行权限认证
	InitPRouter.POST("/get_container", controller.FindAllContController) //开始响应，在我们的控制器文件中调用我的处理函数
}

func DeleteContainerRoute(r *gin.Engine) {
	InitPRouter := r.Group("/admin")
	r.LoadHTMLFiles("./no_right.html")                                //调用中间件，用于进行权限认证
	InitPRouter.POST("/delete_container", controller.DeleteContainer) //开始响应，在我们的控制器文件中调用我的处理函数
}

func AddContainerRoute(r *gin.Engine) {
	InitPRouter := r.Group("/admin")
	r.LoadHTMLFiles("./no_right.html")                  //调用中间件，用于进行权限认证
	InitPRouter.POST("/add_container", controller.Link) //开始响应，在我们的控制器文件中调用我的处理函数
}

func GetServiceRoute(r *gin.Engine) {
	InitPRouter := r.Group("/admin")
	r.LoadHTMLFiles("./no_right.html")                                //调用中间件，用于进行权限认证
	InitPRouter.POST("/get_server", controller.Get_ServiceController) //开始响应，在我们的控制器文件中调用我的处理函数
}

func AddServiceRoute(r *gin.Engine) {
	InitPRouter := r.Group("/admin")
	r.LoadHTMLFiles("./no_right.html")                                //调用中间件，用于进行权限认证
	InitPRouter.POST("/add_server", controller.Add_ServiceController) //开始响应，在我们的控制器文件中调用我的处理函数
}

func DeleteServiceRoute(r *gin.Engine) {
	InitPRouter := r.Group("/admin")
	r.LoadHTMLFiles("./no_right.html")                                      //调用中间件，用于进行权限认证
	InitPRouter.POST("/delete_server", controller.Delete_ServiceController) //开始响应，在我们的控制器文件中调用我的处理函数
}

func UpdateServerRoute(r *gin.Engine) {
	InitPRouter := r.Group("/admin")
	r.LoadHTMLFiles("./no_right.html")                                      //调用中间件，用于进行权限认证
	InitPRouter.POST("/update_server", controller.Update_ServiceController) //开始响应，在我们的控制器文件中调用我的处理函数
}

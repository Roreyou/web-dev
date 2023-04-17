package main

import (
	"back/router"
)

func main() {
	// r := gin.Default()
	/*
		下面是之前的测试代码
		api := r.Group("api") // 创建一个名叫api的路由组,并且该组的统一前缀为/api
		api.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "您已经进入了登陆页面，在这个页面，你可以输入你的个人信息，以便进行_进入后面页面前的_校验"}) //不需要权限校验的页面
		})
	*/

	//调用我们自己的route包中的来创建路由
	r := router.SetupRouter()
	dao.Openmysql()     //打开数据库
	router.Gethtml(r)   //显示登陆页面
	router.CheckUser(r) //登陆判断
	r.Run(":8080")
}

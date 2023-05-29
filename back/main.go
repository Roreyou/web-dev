package main

import (
	"back/dao"
	"back/router"
)

func main() {

	//调用我们自己的route包中的来创建路由
	r := router.SetupRouter()
	dao.Openmysql() //打开数据库
	r.Run(":8081")
}

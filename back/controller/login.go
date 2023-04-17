package controller

import (
	"net/http"

	"back/dao"

	"github.com/gin-gonic/gin"
)

// 存放路由函数的具体处理函数，写路由时调用该函数即可

func LoginInfoController(c *gin.Context) { //处理登陆信息
	//从请求中读取数据
	//获取来自html中的数据,用户名和密码
	name := c.PostForm("username")
	password := c.PostForm("password")
	user := dao.FindaUser(name) //找到用户名相同的用户
	if user.ID == 0 {           //该用户不存在则报错
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"fail": "该用户名不存在",
		})
		return
	}

	if user.Password == password { //密码正确
		c.JSON(http.StatusOK, gin.H{
			"success": "密码正确，成功登陆",
		})
		return
	} else { //密码不正确
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"fail": "密码不正确",
		})
		return
	}
}

func Showhtml(c *gin.Context) { //显示登陆页面
	c.HTML(http.StatusOK, "login.html", nil) //访问/login时呈现出login.html
}

func logon(c *gin.Context) { //注册
	name := c.PostForm("username")
	password := c.PostForm("password")
	user := dao.FindaUser(name) //找到用户名相同的用户
	if user.Username == name {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"fail": "该用户名已存在",
		})
		return
	} else {
		dao.CreateaUser(name, password)
	}
}

func AuthLoginHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user User.Profile
	err := c.ShouldBindJSON(&user)
	if err != nil { //能否获取到正确参数
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	// 校验用户名和密码是否正确
	auser := dao.CreateaUser(user.Username)
	if auser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 2003,
			"msg":  "用户不存在",
		})
		return
	}

	tokenString, _ := middlewares.AuthToken(user.Username)
	c.JSON(http.StatusOK, gin.H{
		"code":     2000,
		"msg":      "success",
		"Token":    tokenString,
		"username": auser.Username,
	})
}

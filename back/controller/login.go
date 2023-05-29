package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"

	"back/dao"

	"github.com/gin-gonic/gin"
	"github.com/langwan/go-jwt-hs256"
)

// 存放路由函数的具体处理函数，写路由时调用该函数即可
func LoginInfoController(c *gin.Context) { //处理登陆信息
	//从请求中读取数据
	//获取来自html中的数据,用户名和密码
	db := dao.Openmysql()
	id := c.PostForm("user_id")
	user_id, _ := strconv.Atoi(id) //要转化成int类型
	password := c.PostForm("user_password")
	hashed := sha256.Sum256([]byte(password))
	hashed_password := hex.EncodeToString(hashed[:])
	fmt.Println(user_id)
	fmt.Println(password)
	user := dao.FindUser(user_id, db) //找到用户名相同的用户

	if user.User_id == 0 { //该用户不存在则报错
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"fail": "该用户名不存在",
		})
		return
	}
	sign, _ := jwt.Sign(user)
	if user.User_password == hashed_password { //密码正确
		c.JSON(http.StatusOK, struct {
			Token string `json:"token"`
		}{Token: sign})
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

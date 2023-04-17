package middlewares

import (
	"back/services"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func AuthToken(username string) gin.HandlerFunc {
	return func(c *gin.Context) { //按照它的要求打了个框架，防止飘红报错
		authHeader := c.Request.Header.Get("authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		mc, ok := services.ParseToken(authHeader)
		if ok == false {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		m := mc.(jwt.MapClaims)
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", m["username"])
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}

}

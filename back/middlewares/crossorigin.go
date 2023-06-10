package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsHander() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		// 用于设置服务器支持的所有跨域请求的方法
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
		// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
		c.Writer.Header().Set("Access-Control-Allow-Headers", "content-type")
		// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			print("进入options")
			fmt.Println(c)
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			//继续处理下面的请求
			c.Next()
		}

	}

}

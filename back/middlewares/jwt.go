package middlewares

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/langwan/go-jwt-hs256"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.RequestURI == "/login" { //处于登录页面
			return
		}
		if token, ok := c.Request.Header["Token"]; ok {
			err := jwt.Verify(token[0])
			if err != nil {
				c.AbortWithStatusJSON(403, "forbidden")
			} else {
				c.AbortWithStatusJSON(403, "forbidden")
			}
		}
	}
}

func getToken(c *gin.Context)(*User,error){
	if token , ok:= c.Request.Header{"Token"}; ok{
		ss := strings.Split(token[0],".")
		at:=&User{}
		payload,err := base64.RawURLEncoding.DecodeString(ss[1])
		if err!= nil{
			return nil,err
		}
		err = json.Unmarshal(payload,at)
		return at,nil
	}
}

// UserInfo 当前用户登录信息
func UserInfo(c *gin.Context) *User {
	token := c.Request.Header.Get("token")
	user := &User{}
	if token == "" {
		return user
	} else {
		claims, err := ParseToken(token)
		if err != nil {
			return user
		} else {
			user.ID = claims.ID
			user.Name = claims.Name
			return user
		}
	}
}

// UserId 当前用户 ID
func UserId(c *gin.Context) int64 {
	return UserInfo(c).ID
}

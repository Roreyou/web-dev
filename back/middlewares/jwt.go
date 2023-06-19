package middlewares

import (
	"encoding/base64"
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

func GetToken(c *gin.Context) (*User, error) {
	if token1, ok := c.Request.Header["Token"]; ok {
		ss := strings.Split(token1[0], ".")
		at := &User{}
		payload, err := base64.RawURLEncoding.DecodeString(ss[1])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(payload, at)
		return at, nil
	} else {
		return nil, gin.Error{}
	}
}

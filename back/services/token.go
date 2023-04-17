package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// JwtClaims 创建自己的Claims
type MyClaims struct {
	jwt.StandardClaims
	Username string
}

var (
	//盐
	secret                 = []byte("wondersafebox")                   // 后续加密增加盐增加复杂度
	TokenExpired     error = errors.New("Token is expired")            // token错误类型提炼
	TokenNotValidYet error = errors.New("Token not active yet")        // token错误类型提炼
	TokenMalformed   error = errors.New("That's not even a token")     // token错误类型提炼
	TokenInvalid     error = errors.New("Couldn't handle this token:") // token错误类型提炼
)

const TokenExpireDuration = time.Hour * 24 //设置过期时间

var Secret = []byte("secret") //密码自行设定

func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "superxon",                                 // 签发人
		},
		username, // 自定义字段
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Secret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

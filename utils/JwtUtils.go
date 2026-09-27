package utils

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 自定义声明
type CustomClaims struct {
	UserId int64 `json:"userid"`
	Role   int8  `json:"role"`
	jwt.RegisteredClaims
}

// 生成token
func GenerateToken(UserId int64, Role int8) (string, error) {
	fmt.Println(MyConfig)
	slog.Info("Token存活时间", "time", MyConfig.Jwt.ExpireDuration)

	claim := CustomClaims{
		UserId: UserId,
		Role:   Role,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(MyConfig.Jwt.ExpireDuration)),
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 签发者
			Issuer: MyConfig.Jwt.Issur,
			// 不晚于
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   "user_token",
		},
	}
	// token包含jwt的头部和负载部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// SignedString需要传入[]byte
	return token.SignedString([]byte(MyConfig.Jwt.SecretKey))
}

// 解析token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (any, error) {
		// t.Method.(*jwt.SigningMethodHMAC)
		// 这是一个类型断言，断言t.Method是*jwt.SigningMethodHMAC类型
		// *jwt.SigningMethodHMAC实现了SigningMethod接口
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("加密算法错误")
		}
		return []byte(MyConfig.Jwt.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("非法的token")
}

// 获取用户id
func GetUserId(ctx *gin.Context) int64 {
	id, exist := ctx.Get("userId")
	if !exist {
		// 不存在
		panic("用户id不存在")
	}
	switch id64 := id.(type) {
	case int64:
		return id64
	default:
		panic("用户id类型错误")
	}
}

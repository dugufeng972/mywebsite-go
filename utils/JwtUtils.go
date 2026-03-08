package utils

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/yaml.v3"
)

// 自定义声明
type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// jwt配置
type JWTConfig struct {
	SecretKey      string        `yaml:"secretkey"`      // 密钥
	Issur          string        `yaml:"issur"`          // 签发者
	ExpireDuration time.Duration `yaml:"expireduration"` // 存活时间
}

type JWTUtil struct {
	Config *JWTConfig `yaml:"jwtconfig"`
}

// 初始化一个JWTUtil
func NewJWTUtil() *JWTUtil {
	// 运行的main，所以文件位置要相对于main来说
	byteStr, err := os.ReadFile("./application.yml")
	if err != nil {
		panic("读取配置文件出异常了")
	}
	jwtUtil := JWTUtil{}
	err = yaml.Unmarshal(byteStr, &jwtUtil)
	if err != nil {
		panic("解析配置出异常了")
	}
	slog.Info("jwtutil内容", "内容", jwtUtil)
	if jwtUtil.Config == nil {
		panic("解析jwt配置失败")
	}
	return &jwtUtil
}

// 生成token
func (j *JWTUtil) GenerateToken(username string) (string, error) {
	// slog.Info(string(j.config.ExpireDuration))
	fmt.Println(j.Config.ExpireDuration)
	claim := CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.Config.ExpireDuration)),
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 签发者
			Issuer: j.Config.Issur,
			// 不晚于
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   "user_token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(j.Config.SecretKey))
}

// 解析token
func (j *JWTUtil) ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (any, error) {
		// t.Method.(*jwt.SigningMethodHMAC)
		// 这是一个类型断言，断言t.Method是*jwt.SigningMethodHMAC类型
		// *jwt.SigningMethodHMAC实现了SigningMethod接口
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("加密算法错误")
		}
		return []byte(j.Config.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("非法的token")
}

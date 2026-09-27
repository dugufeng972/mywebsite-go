// 工具模块
package utils

import (
	"log/slog"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type JWTConfig struct {
	SecretKey      string        `yaml:"secretkey"`      // 密钥
	Issur          string        `yaml:"issur"`          // 签发者
	ExpireDuration time.Duration `yaml:"expireduration"` // 存活时间
}

type AuthConfig struct {
	Path []string `yaml:"path"`
}

// func (a AuthConfig) IncludePath(path string) bool {
// 	slices.Contains(a.Path, path)
// 	for _, p := range a.Path {
// 		if p == path {
// 			return true
// 		}
// 	}
// 	return false
// }

type Config struct {
	Jwt  JWTConfig  `yaml:"jwtconfig"`
	Auth AuthConfig `yaml:"auth"`
}

// 全局配置
var MyConfig = Config{}

func Init() {
	byteStr, err := os.ReadFile("./application.yml")
	if err != nil {
		panic("读取配置失败")
	}
	err = yaml.Unmarshal(byteStr, &MyConfig)
	if err != nil {
		slog.Error("解析配置出异常了")
		panic("解析配置出异常了")
	}
	slog.Info("配置读取成功")
}

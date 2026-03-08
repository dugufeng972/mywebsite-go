package utils

import (
	"github.com/gin-gonic/gin"
)

// 错误处理函数
func MidlErrorHandle(err error, msg string, ctx *gin.Context) bool {
	if err != nil {
		ctx.JSON(500, gin.H{
			"msg": msg,
		})
		return true
	}
	return false
}

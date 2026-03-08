package utils

import "github.com/gin-gonic/gin"

func Result(ctx *gin.Context, code int, data any) {
	ctx.JSON(code, data)
}

func SuccessResp(ctx *gin.Context, code int, data any) {
	ctx.JSON(code, data)
}

func ErrorResp(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"msg": msg,
	})
}

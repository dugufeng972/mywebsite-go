package user

import (
	"webgo/constants"

	"github.com/gin-gonic/gin"
)

func LoginAuth(ctx *gin.Context) {
	_, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(401, gin.H{
			"message": "用户未登录",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "用户已登录",
	})
}

// 获取验证码
func GetAuthCode(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		constants.MsgStr: constants.Success,
		"data":           4567,
	})
}

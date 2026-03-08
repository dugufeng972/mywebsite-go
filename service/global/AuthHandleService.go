package global

import (
	"webgo/utils"

	"github.com/gin-gonic/gin"
)

// 权限验证
func AuthHandleMid(ctx *gin.Context) {
	if ctx.FullPath() == "/user/login" || ctx.FullPath() == "/user/register" {
		ctx.Next()
		return
	}
	token := ctx.GetHeader("cookie")
	if token == "" {
		panic("没有token")
	}
	claims, err := utils.NewJWTUtil().ParseToken(token)
	if err != nil {
		panic(err.Error())
	}
	ctx.Set("username", claims.Username)
	ctx.Next()
}

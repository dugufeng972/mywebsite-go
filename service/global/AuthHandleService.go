package global

import (
	"log/slog"
	"slices"
	"webgo/utils"

	"github.com/gin-gonic/gin"
)

// 权限验证
func AuthHandleMid(ctx *gin.Context) {
	slog.Info("请求路径", "路径", ctx.FullPath())
	if slices.Contains(utils.MyConfig.Auth.Path, ctx.FullPath()) {
		ctx.Next()
		return
	}
	token, cookieErr := ctx.Cookie("token")
	if cookieErr != nil {
		slog.Error("权限验证", "info", "获取token失败", "error msg", cookieErr)
		ctx.JSON(401, gin.H{
			"message": "用户未登录",
		})
		return
	}
	if token == "" {
		panic("没有token")
	}
	claims, err := utils.ParseToken(token)
	if err != nil {
		panic(err.Error())
	}
	slog.Info("token验证成功", "用户id", claims.UserId)
	ctx.Set("userId", claims.UserId)
	ctx.Next()
}

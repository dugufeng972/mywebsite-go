package global

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 全局中间件 -- 异常处理

func ErrorHandleMid(ctx *gin.Context) {
	defer func(ctx *gin.Context) {
		err := recover()
		if err != nil {
			slog.Warn("全局异常处理中间件处理", "异常信息", err.(string))
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "服务器开小差了",
			})
		}
	}(ctx)

	ctx.Next()
}

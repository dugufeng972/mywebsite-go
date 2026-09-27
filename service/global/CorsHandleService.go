package global

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsHandleMiddle(ctx *gin.Context) {
	method := ctx.Request.Method
	origin := ctx.Request.Header.Get("Origin")
	// 相当于设置为 *
	if origin != "" {
		// 设置允许任何地址的请求
		ctx.Header("Access-Control-Allow-Origin", origin)
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		// 允许使用的请求头
		ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		// 允许客户端访问的响应头
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		// 允许客户端携带cookie
		ctx.Header("Access-Control-Allow-Credentials", "true")
	}

	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}
	ctx.Next()
}

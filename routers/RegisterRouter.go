package routers

import (
	"webgo/routers/user"

	"github.com/gin-gonic/gin"
)

// 集中注册路由函数
func RegisterRouters(r *gin.Engine) {
	user.UserRouter(r)
}

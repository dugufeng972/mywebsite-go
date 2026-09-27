package routers

import (
	"webgo/routers/book"
	"webgo/routers/user"

	"github.com/gin-gonic/gin"
)

// 集中注册路由
func RegisterRouters(r *gin.Engine) {
	user.UserRouter(r)
	book.BookRouter(r)
}

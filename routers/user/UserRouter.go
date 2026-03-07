package user

import (
	"webgo/service/user"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/login", user.UserServiceLogin)
	}
}

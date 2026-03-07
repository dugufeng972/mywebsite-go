// Package user
package user

import (
	"fmt"
	"webgo/domain/po"
	"webgo/domain/vo"
	"webgo/models"
	"webgo/utils"

	"github.com/gin-gonic/gin"
)

func UserServiceLogin(ctx *gin.Context) {
	user := vo.UserLoginVO{}
	err := ctx.ShouldBindJSON(&user)
	utils.MidlErrorHandle(err, "数据绑定错误", ctx)
	fmt.Println(user.UserName)
	u := po.UserPO{}
	models.DB.Table("swuser").Where("username=?", user.UserName).Find(&u)
	// models.DB.Raw("select * from swuser where username=?", user.UserName).Scan(&u)
	ctx.JSON(200, gin.H{
		"msg": u.UserName + " " + u.PassWord,
	})
	fmt.Println(u)
}

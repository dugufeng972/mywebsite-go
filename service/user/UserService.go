// Package user
package user

import (
	"log/slog"
	"time"
	"webgo/domain/po"
	"webgo/domain/vo"
	"webgo/models"
	"webgo/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// 登录中间件
func UserServiceLogin(ctx *gin.Context) {
	user := vo.UserLoginVO{}
	// err := ctx.ShouldBindJSON(&user)
	if utils.MidlErrorHandle(ctx.ShouldBindJSON(&user), "数据绑定错误", ctx) {
		return
	}
	u := po.UserPO{}
	models.DB.Table("swuser").Where("username=?", user.UserName).Find(&u)
	if u.UserName == "" {
		utils.Result(ctx, 400, "用户不存在")
		return
	}
	jwtUtil := utils.NewJWTUtil()
	token, _ := jwtUtil.GenerateToken(u.UserName)
	utils.Result(ctx, 200, token)
}

// 注册中间件
func UserServiceRegister(ctx *gin.Context) {
	user := vo.UserLoginVO{}
	if utils.MidlErrorHandle(ctx.ShouldBindJSON(&user), "数据绑定失败", ctx) {
		return
	}
	var count int64 = 0
	res := models.DB.Model(&po.Swuser{}).Where("username=?", user.UserName).Count(&count)
	if res.Error != nil {
		utils.Result(ctx, 500, "异常")
		slog.Error(res.Error.Error())
		return
	}
	if count != 0 {
		utils.Result(ctx, 200, "用户名存在")
		return
	}
	userPO := po.Swuser{}
	if copier.Copy(&userPO, user) != nil {
		panic("数据copy失败")
	}
	userPO.ID = time.Now().UnixMilli()
	userPO.Type = 1
	userPO.Status = 1
	userPO.CreateTime = time.Now()
	userPO.UpdateTime = time.Now()
	userPO.Creater = new(int32)
	*userPO.Creater = 1
	userPO.Updater = new(int32)
	*userPO.Updater = 1
	result := models.DB.Table("swuser").Create(userPO)
	if result.Error != nil || result.RowsAffected == 0 {
		utils.Result(ctx, 200, "注册失败")
	}
	token, err := utils.NewJWTUtil().GenerateToken(user.UserName)
	if err != nil {
		utils.Result(ctx, 200, "注册失败")
	}
	utils.Result(ctx, 200, token)
}

func UserTest(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	ctx.JSON(200, gin.H{
		"msg":  username,
		"test": "test",
	})
}

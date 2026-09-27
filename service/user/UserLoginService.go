// Package user
package user

import (
	"log/slog"
	"time"
	"webgo/constants"
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
	u := po.Swuser{}
	models.DB.Table("swuser").Where("username=?", user.UserName).Find(&u)
	if u.Username == "" {
		utils.Result(ctx, 400, "用户不存在")
		return
	}
	// jwtUtil := utils.NewJWTUtil()
	token, _ := utils.GenerateToken(u.ID, int8(u.Type))
	ctx.SetCookie("token", token, 604800, "/", "localhost", false, true)
	ctx.JSON(200, gin.H{
		"message": "登录成功",
		"data": map[string]any{
			"token": token,
			"info": vo.UserInfoVO{
				UserName:  u.Username,
				Role:      int8(u.Type),
				Stauts:    int8(u.Status),
				AvatarURL: "",
			},
		},
	})
}

// 注册中间件
func UserServiceRegister(ctx *gin.Context) {
	user := vo.UserLoginVO{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		slog.Error("数据绑定失败")
		ctx.JSON(200, gin.H{
			"message": "前端数据异常",
		})
		return
	}
	if user.Code != "4567" {
		ctx.JSON(200, gin.H{
			"message": "验证码错误",
		})
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
		utils.Result(ctx, 200, constants.CodeMessages[constants.USERNAMEEXIEST])
		return
	}
	userPO := po.Swuser{}
	if copier.Copy(&userPO, user) != nil {
		panic("数据copy失败")
	}
	userPO.ID = models.Node.Generate().Int64()
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
		slog.Info("用户注册", "msg", "数据库写入失败")
		utils.Result(ctx, 200, "注册失败")
	}
	token, err := utils.GenerateToken(userPO.ID, int8(userPO.Type))
	if err != nil {
		slog.Info("用户注册", "msg", "jwt token生成失败")
		utils.Result(ctx, 200, "注册失败")
		return
	}
	// 设置 Cookie 有效期为 7 天
	ctx.SetCookie("token", token, 604800, "/", "localhost", false, true)
	ctx.JSON(200, gin.H{
		"message": "注册成功",
		"data": map[string]any{
			"token": token,
			"info": vo.UserInfoVO{
				UserName:  userPO.Username,
				Role:      int8(userPO.Type),
				Stauts:    int8(userPO.Status),
				AvatarURL: "",
			},
		},
	})
}

// 登录验证接口

// 这是一个测试用接口
func UserTest(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	ctx.JSON(200, gin.H{
		"msg":  username,
		"test": "test",
	})
}

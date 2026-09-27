// 用户相关的服务
package user

import (
	"webgo/domain/po"
	"webgo/domain/vo"
	"webgo/models"
	"webgo/utils"

	"github.com/gin-gonic/gin"
)

// 查询用户信息
func GetUserInfo(ctx *gin.Context) {
	userId := utils.GetUserId(ctx)
	userPO := po.Swuser{}
	models.DB.Where("id=?", userId).Find(&userPO)
	userInfoVO := vo.UserInfoVO{}
	userInfoVO.UserName = userPO.Username
	userInfoVO.Role = int8(userPO.Type)
	userInfoVO.AvatarURL = *userPO.AvatarURL
	userInfoVO.Stauts = int8(userPO.Status)
	ctx.JSON(200, gin.H{
		"message": "用户信息查询成功",
		"data":    userInfoVO,
	})
}

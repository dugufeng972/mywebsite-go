package book

import (
	"log/slog"
	"webgo/domain/vo"
	"webgo/utils"

	"github.com/gin-gonic/gin"
)

func UploadBook(ctx *gin.Context) {
	slog.Info("upload book")
	formData := vo.BookInfoVO{}
	if err := ctx.ShouldBind(&formData); err != nil {
		slog.Error("book上传", "info", "vo数据绑定失败")
		ctx.JSON(200, gin.H{
			"message": "上传的数据异常，不符合后端要求",
		})
		return
	}
	slog.Info(formData.File.Filename)
	// fileNameList := strings.Split(formData.File.Filename, ".")
	// if len(fileNameList) <= 1 {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "没有后缀名或者没有文件名",
	// 	})
	// 	return
	// }

	// 把book放进数据库
	// 1. 根据名字和上传者id查询该book是否在数据库

	// 1.1 存在，更新book表相关字段，向book_chapters表追加正文(后面统一操作)

	// 1.2 不存在，在book表中创建该book的行，向book_chapters表写入正文(后面统一操作)

	// 2. 向book_chapters表写入数据
	if file, err := formData.File.Open(); err == nil {
		utils.TxtParserAndSave(&file)
		defer func() {
			err := file.Close()
			if err != nil {
				file = nil
				panic("文件释放失败")
			}
			slog.Info("uploadbook", "info", "文件句柄已释放")
		}()
	}
	// userId := utils.GetUserId(ctx)
	// 保存文件
	// filePath := "./static/book/" + strconv.FormatInt(userId, 10) + "/" + fileNameList[0] + "/" + formData.File.Filename
	// ctx.SaveUploadedFile(formData.File, filePath)
	ctx.JSON(200, gin.H{
		"message": "upload success",
	})
}

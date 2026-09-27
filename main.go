package main

import (
	"log/slog"
	"webgo/models"
	"webgo/routers"
	"webgo/service/global"
	"webgo/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 初始化配置
	utils.Init()
	// 初始化数据库
	models.InitDB()
	// 加载静态文件目录
	r.Static("/static", "./static")
	// 注册全局中间件
	r.Use(global.CorsHandleMiddle, global.ErrorHandleMid, global.AuthHandleMid)
	// r.Use(global.ErrorHandleMid)
	// r.Use(global.AuthHandleMid)
	routers.RegisterRouters(r)
	slog.Info("正在启动服务：监听8080端口")
	err := r.Run()
	if err != nil {
		slog.Error("服务启动失败", "info", err)
		return
	}
}

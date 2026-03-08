package main

import (
	"webgo/models"
	"webgo/routers"
	"webgo/service/global"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.InitDB()
	r.Use(global.ErrorHandleMid)
	r.Use(global.AuthHandleMid)
	routers.RegisterRouters(r)
	r.Run()
}

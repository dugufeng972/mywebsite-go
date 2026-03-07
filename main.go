package main

import (
	"webgo/models"
	"webgo/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.InitDB()
	routers.RegisterRouters(r)
	r.Run()
}

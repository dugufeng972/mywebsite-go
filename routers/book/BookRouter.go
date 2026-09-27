package book

import (
	"webgo/service/book"

	"github.com/gin-gonic/gin"
)

func BookRouter(r *gin.Engine) {
	bookRouter := r.Group("/book")
	{
		bookRouter.POST("/upbook", book.UploadBook)
	}
}

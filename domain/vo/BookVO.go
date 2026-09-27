package vo

import "mime/multipart"

// book上传
type BookInfoVO struct {
	BookName     string                `form:"bookname"`
	Gender       string                `form:"gender"`
	Introduction string                `form:"intro"`
	File         *multipart.FileHeader `form:"booktxt"`
	Cover        *multipart.FileHeader `form:"cover"`
	Tags         []string              `form:"tags"`
}

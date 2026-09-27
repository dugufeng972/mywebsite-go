// Package constants 用于定义错误码和对应的错误信息
package constants

// 错误码
const (
	USERNAMEEXIEST = 400 + iota
)

// 错误信息
var CodeMessages = map[int]string{
	USERNAMEEXIEST: "用户名已存在",
}

// 固定字段
const (
	MsgStr  = "message"
	Success = "请求成功"
)

package vo

// 用户登录的vo
type UserLoginVO struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Vid      string `json:"vid"`
	Code     string `json:"code"`
}

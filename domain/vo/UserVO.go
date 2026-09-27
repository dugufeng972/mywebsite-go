package vo

// 用户登录和注册的vo
type UserLoginVO struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	Code     string `json:"code"`
}

// 用户信息的vo
type UserInfoVO struct {
	UserName  string `json:"username"`
	Role      int8   `json:"role"`
	Stauts    int8   `json:"status"`
	AvatarURL string `json:"avatar"`
}

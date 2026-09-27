package po

type UserPO struct {
	UserId    int64  `gorm:"column:id" json:"userId"`
	UserName  string `gorm:"column:username" json:"username"`
	PassWord  string `gorm:"column:password" json:"password"`
	UserRole  int8   `gorm:"column:type" json:"role"`
	Status    int8   `gorm:"column:status" json:"status"`
	AvatarUrl string `gorm:"column:avatar_url" json:"avatar"`
}

func (UserPO) TableName() string {
	return "swuser"
}

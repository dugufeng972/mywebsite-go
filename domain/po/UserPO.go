package po

type UserPO struct {
	UserName string `gorm:"column:username" json:"username"`
	PassWord string `gorm:"column:password" json:"password"`
}

func (UserPO) TableName() string {
	return "swuser"
}

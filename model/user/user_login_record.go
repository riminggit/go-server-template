package userModel

type UserLoginRecord struct {
	ID       int    `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	UserId   int    `json:"user_id" gorm:"type:bigint;not null"`                    // 用户id
	LoginAt  int64  `json:"login_at"`
	ComeFrom string `json:"come_from" gorm:"type:varchar(20);"` // 登陆途径
	CreateAt int64  `json:"create_at"`                          // 创建时间
	DeleteAt int64  `json:"delete_at"`
	UpdateAt int64  `json:"update_at"`
	IsUse    int    `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

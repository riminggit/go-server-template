package userModel

import (
	"time"
)

type UserLoginRecord struct {
	ID       int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	UserId   int       `json:"user_id" gorm:"type:bigint;not null"`                   // 用户id
	LoginAt  time.Time `json:"login_at"`
	ComeFrom string    `json:"come_from" gorm:"type:varchar(20);"` // 登陆途径
	CreateAt time.Time `json:"create_at"`                          // 创建时间
	DeleteAt time.Time `json:"delete_at"`
	UpdateAt time.Time `json:"update_at"`
	IsUse    int       `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

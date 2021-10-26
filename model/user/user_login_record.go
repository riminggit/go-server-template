package userModel

import (
	"time"
)

type UserLoginRecord struct {
	ID       int `gorm:"autoIncrement;primaryKey;type:int(11);uniqueIndex;not null"` // id
	UserId   int `gorm:"type:int(11);not null"`                                      // 用户id
	LoginAt  time.Time
	ComeFrom string `gorm:"type:varchar(20);"`    // 登陆途径
	IsUse    int    `gorm:"type:int(3);not null"` // 是否删除:0 删除 1未删除
}

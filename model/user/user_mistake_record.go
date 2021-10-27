package userModel

import (
	"time"
)

type UserMistakeRecord struct {
	ID         int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	UserId     int       `gorm:"type:int(11);not null"`                                  // 用户id
	TopicId    int       `gorm:"type:int(11);not null"`                                  // 对应题目id
	MistakeNum int       `gorm:"type:int(3);not null"`                                   // 错误次数
	RightNum   int       `gorm:"type:int(3);not null"`                                   // 正确次数
	CreateAt   time.Time // 创建时间
	UpdataAt   time.Time // 更新时间
	IsUse      int       `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

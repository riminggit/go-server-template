package userModel

import (
	"time"
)

type UserExperience struct {
	ID         int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	UserId     int       `gorm:"type:int(11);not null"`                                  // 用户id
	Experience int       `gorm:"type:int(11);not null"`                                  // 用户经验
	Level      int       `gorm:"type:int(11);not null"`                                  // 用户等级
	IsTest     int       `gorm:"type:int(3);not null"`                                   // 是否考试 0 否 1 是
	CanTest    int       `gorm:"type:int(3);not null"`                                   // 是否可以参加考试  0 否 1 是
	UpdateAt   time.Time // 修改时间
	IsUse      int       `gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

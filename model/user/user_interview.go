package userModel

import (
	"time"
)

type UserInterview struct {
	ID                int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	UserId            int       `gorm:"type:int(11);not null"`                                  // 用户id
	CompanyName       string    `gorm:"type:varchar(50);"`                                      // 公司名
	InterviewTime     time.Time // 面试时间
	InterviewSchedule string    `gorm:"type:varchar(50);"` // 面试阶段
	InterviewStatus   int       `gorm:"type:int(3);"`      // 0 未开始 1进行中 2已完成
	InterviewResult   string    `gorm:"type:varchar(50);"` // 面试结果
	CreateAt          time.Time // 创建时间
	UpdataAt          time.Time // 更新时间
	IsUse             int       `gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

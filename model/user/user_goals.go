package userModel

import (
	"time"
)

type UserGoals struct {
	ID             int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	UserId         int       `gorm:"type:bigint;not null"`                                   // 用户id
	GoalType       int       `gorm:"type:int(3);"`                                           // 1 是自定义目标  2是模块目标(classify) 3是知识点目标(type)   4是自选题目目标
	CustomGoals    string    `gorm:"type:text;"`                                             // 自定义目标内容
	ClassifyIdList string    `gorm:"type:text;"`                                             // 方向目标id 数组
	TypeIdList     string    `gorm:"type:text;"`                                             // 模块目标id 数组
	TopicIdList    string    `gorm:"type:text;"`                                             // 题目id 数组
	CreateAt       time.Time `json:"create_at"`                                              // 创建时间
	DeleteAt       time.Time `json:"delete_at"`
	UpdateAt       time.Time `json:"update_at"`
	StartAt        time.Time // 目标设定开始时间
	EndAt          time.Time // 目标设定结束时间
	Progress       int       `gorm:"type:int(3);not null"`           // 进度，最大100
	GoalsStatus    int       `gorm:"type:int(3);"`                   // 目标是 1 进行中   2已放弃(未完成)  3已完成
	IsUse          int       `gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

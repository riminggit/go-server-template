package userModel

import (
	"time"
)

type UserTestRecord struct {
	ID             int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	UserId         int       `gorm:"type:bigint;not null"`                                   // 用户id
	Source         int       `gorm:"type:int(3);"`                                           // 考试成绩
	TestAt         time.Time // 考试开始时间
	TestEnd        time.Time // 考试结束时间
	TestTopicList  string    `gorm:"type:text;"`   // 考试题目列表数组
	ClassifyIdList string    `gorm:"type:text;"`   // 涉及的classify数组
	TypeIdList     string    `gorm:"type:text;"`   // 涉及的type数组
	CorrectNum     int       `gorm:"type:int(3);"` // 正确数量
	MistakeNum     int       `gorm:"type:int(3);"` // 错误数量
	CreateAt       time.Time `json:"create_at"`    // 创建时间
	DeleteAt       time.Time `json:"delete_at"`
	UpdateAt       time.Time `json:"update_at"`
	IsUse          int       `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

package userModel

import (
	"time"
)

type UserTestRecord struct {
	ID             int       `gorm:"autoIncrement;primaryKey;type:int(11);uniqueIndex;not null"` // id
	UserId         int       `gorm:"type:int(11);not null"`                                      // 用户id
	Source         int       `gorm:"type:int(3);"`                                               // 考试成绩
	TestAt         time.Time // 考试开始时间
	TestEnd        time.Time // 考试结束时间
	TestTopicList  string    `gorm:"type:varchar(255);"` // 考试题目列表数组
	ClassifyIdList string    `gorm:"type:varchar(255);"` // 涉及的classify数组
	TypeIdList     string    `gorm:"type:varchar(255);"` // 涉及的type数组
	CorrectNum     int       `gorm:"type:int(3);"`       // 正确数量
	MistakeNum     int       `gorm:"type:int(3);"`       // 错误数量
	CreateAt       time.Time // 创建时间
	IsUse          int       `gorm:"type:int(3);not null"` // 是否删除:0 删除 1未删除
}

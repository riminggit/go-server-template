package userModel

import (
	"time"
)

type UserFeedback struct {
	ID                 int       `gorm:"autoIncrement;primaryKey;type:int(11);uniqueIndex;not null"` // id
	UserId             int       `gorm:"type:int(11);not null"`                                      // 用户id
	FeedbackType       int       `gorm:"type:int(3);not null"`                                       // 反馈类型 1 投诉  2 建议  3 题目相关
	FacebackAnswer     string    `gorm:"type:varchar(255);"`                                         // 反馈回复
	TopicId            int       `gorm:"type:int(11);"`                                              // 对应题目id
	Content            string    `gorm:"type:text;"`                                                 // 反馈内容
	CreateAt           time.Time // 创建时间
	FacebackAnswerTime time.Time // 反馈时间
	Grade              int       `gorm:"type:int(3);"`         // 评分
	IsUse              int       `gorm:"type:int(3);not null"` // 是否删除:0 删除 1未删除
}

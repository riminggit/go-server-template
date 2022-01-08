package userModel

import (
	"time"
)

type UserFeedback struct {
	ID                 int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	UserId             int       `json:"user_id" gorm:"type:int(11);not null"`                                  // 用户id
	FeedbackType       int       `json:"feedback_ype" gorm:"type:int(3);not null"`                                   // 反馈类型 1 投诉  2 建议  3 题目相关
	FacebackAnswer     string    `json:"faceback_answer" gorm:"type:varchar(255);"`                                     // 反馈回复
	TopicId            int       `json:"topic_id" gorm:"type:int(11);"`                                          // 对应题目id
	Content            string    `json:"content" gorm:"type:text;"`                                             // 反馈内容
	CreateAt           time.Time `json:"create_at"`                                              // 创建时间
	DeleteAt           time.Time `json:"delete_at"`
	UpdateAt           time.Time `json:"update_at"`
	FacebackAnswerTime time.Time // 反馈时间
	Grade              int       `json:"grade" gorm:"type:int(3);"`                   // 评分
	IsRead			   int 		 `json:"is_read" gorm:"type:int(3);"`                   // 评分
	IsUse              int       `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

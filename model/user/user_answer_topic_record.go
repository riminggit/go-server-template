package userModel

import (
	"time"
)

type UserAnswerTopicRecord struct {
	ID               int       `gorm:"autoIncrement;primaryKey;type:int(11);uniqueIndex;not null"` // id
	UserId           int       `gorm:"type:int(11);not null"`                                      // 用户id
	TopicIdList      string    `gorm:"type:varchar(255);"`                                         // 题目id列表
	AnswerStart      time.Time // 答题开始时间
	AnswerEnd        time.Time // 答题结束时间
	AnswerNum        int       `gorm:"type:int(3);"`         // 答题数量
	AnswerCorrectNum int       `gorm:"type:int(3);"`         // 答题正确数量
	IsAchieve        int       `gorm:"type:int(3);"`         // 是否完成答题 0否  1 完成
	TopicDifficulty  int       `gorm:"type:int(3);"`         // 题目难度，由所有的题目难度指数相加后除题目数
	TopicLevel       int       `gorm:"type:int(3);"`         // 题目等级，由所有的题目等级指数相加后除题目数
	Remark           string    `gorm:"type:varchar(255);"`   // 备注
	IsUse            int       `gorm:"type:int(3);not null"` // 是否删除:0 删除 1未删除
}

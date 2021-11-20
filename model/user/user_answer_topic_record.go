package userModel

import (
	"time"
)

type UserAnswerTopicRecord struct {
	ID               int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	UserId           int       `json:"user_id" gorm:"type:int(11);not null"`                   // 用户id
	TopicIdList      string    `json:"topic_id_list" gorm:"type:varchar(100);"`                // 题目id列表
	TopicSetId       string    `json:"topic_set_id" gorm:"type:int(11);"`                      // 套题id
	AnswerStart      time.Time `json:"answer_start"`                                           // 答题开始时间
	AnswerEnd        time.Time `json:"answer_end"`                                             // 答题结束时间
	AnswerNum        int       `json:"answer_num" gorm:"type:int(3);"`                         // 答题数量
	AnswerCorrectNum int       `json:"answer_correct_num" gorm:"type:int(3);"`                 // 答题正确数量
	IsAchieve        int       `json:"is_achieve" gorm:"type:int(3);"`                         // 是否完成答题 0否  1 完成
	TopicDifficulty  int       `json:"topic_difficulty" gorm:"type:int(3);"`                   // 题目难度，由所有的题目难度指数相加后除题目数
	TopicLevel       int       `json:"topic_level" gorm:"type:int(3);"`                        // 题目等级，由所有的题目等级指数相加后除题目数
	Remark           string    `json:"remark" gorm:"type:varchar(255);"`                       // 备注
	CreateAt         time.Time `json:"create_at"`                                              // 创建时间
	DeleteAt         time.Time `json:"delete_at"`
	UpdateAt         time.Time `json:"update_at"`
	IsUse            int       `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

package userModel

import (
	"time"
)

// 用redis存，然后查的时候直接查redis，同步更新到这个表里面，如果有数据就是已读，没有就是未读
type UserTopicRead struct {
	ID               int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	UserId           int       `json:"user_id" gorm:"type:int(11);not null"`                   // 用户id
	TopicId          int       `json:"topic_id" gorm:"type:int(11);not null"`                  // 对应题目id
	IsRead           int       `json:"is_read" gorm:"type:int(3);"`                            // 是否阅读 0 未读  1已读
	ReadNum          int       `json:"read_num" gorm:"type:int(4);"`                           // 阅读次数
	Degree           int       `json:"degree" gorm:"type:int(3);"`                             // 难度，简单：0，中等：1，难：2，极难：3
	Level            int       `json:"level" gorm:"type:int(3);"`                              // 1 初级 ， 2 中级 ， 3 高级 ， 4 资深 ，5 专家 ， 6 资深专家 ， 7 研究员
	IsBaseTopic      int       `json:"is_base_topic" gorm:"type:int(3);"`                      // 0 否  1是 ，是否是基础题
	IsImportantTopic int       `json:"is_important_topic" gorm:"type:int(3);"`                 // 0 否  1是 ，是否是重点题
	CreateAt         time.Time `json:"create_at"`                                              // 创建时间
	DeleteAt         time.Time `json:"delete_at"`
	UpdateAt         time.Time `json:"update_at"`
}

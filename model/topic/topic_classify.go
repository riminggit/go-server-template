package topicModel

import "time"

type TopicClassify struct {
	ID         int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	TopicId    int       `gorm:"type:int(11);not null"`                                  // 对应题目id
	ClassifyId int       `gorm:"type:int(11);not null"`                                  // 类型id
	CreateAt   time.Time `json:"create_at"`                                              // 创建时间
	DeleteAt   time.Time `json:"delete_at"`
	UpdateAt   time.Time `json:"update_at"`
}

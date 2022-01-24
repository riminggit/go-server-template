package topicModel

import "time"

type TopicTag struct {
	ID       int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	TopicId  int       `gorm:"type:bigint;not null"`                                  // 对应题目id
	TagId    int       `gorm:"type:bigint;not null"`                                  // 标签id
	CreateAt time.Time `json:"create_at"`                                              // 创建时间
	DeleteAt time.Time `json:"delete_at"`
	UpdateAt time.Time `json:"update_at"`
}

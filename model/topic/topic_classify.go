package topicModel

type TopicClassify struct {
	ID         int   `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	TopicId    int   `gorm:"type:bigint;not null"`                                   // 对应题目id
	ClassifyId int   `gorm:"type:bigint;not null"`                                   // 类型id
	CreateAt   int64 `json:"create_at"`                                              // 创建时间
	DeleteAt   int64 `json:"delete_at"`
	UpdateAt   int64 `json:"update_at"`
}

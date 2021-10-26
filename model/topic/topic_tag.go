package topicModel

type TopicTag struct {
	ID      int `gorm:"autoIncrement;primaryKey;type:int(11);uniqueIndex;not null"` // id
	TopicId int `gorm:"type:int(11);not null"`                                      // 对应题目id
	TagId   int `gorm:"type:int(11);not null"`                                      // 标签id
	IsUse   int `gorm:"type:int(3);not null"`                                       // 是否删除:0 删除 1未删除
}

package topicModel

type TopicType struct {
	ID      int `gorm:"autoIncrement;primaryKey;type:int(11);uniqueIndex;not null"` // id
	TopicId int `gorm:"type:int(11);not null"`                                      // 对应题目id
	TypeId  int `gorm:"type:int(11);not null"`                                      // 分类id
	IsUse   int `gorm:"type:int(3);not null"`                                       // 是否删除:0 删除 1未删除
}

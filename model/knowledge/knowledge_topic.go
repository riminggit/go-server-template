package knowledgeModel

type KnowledgeTopic struct {
	ID          int   `json:"id" gorm:" column:id;AUTO_INCREMENT;comment:id;not null"` // id
	KnowledgeId int   `json:"knowledge_id" gorm:"type:bigint;not null"`                // 知识点id
	TopicId     int   `json:"topic_id" gorm:"type:bigint;not null"`                    // 题目id
	CreateAt    int64 `json:"create_at"`                                               // 创建时间
	DeleteAt    int64 `json:"delete_at"`
	UpdateAt    int64 `json:"update_at"`
	IsUse       int   `gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

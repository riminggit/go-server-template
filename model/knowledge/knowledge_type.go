package knowledgeModel

import (
	"time"
)

type KnowledgeType struct {
	ID         int       `json:"id" gorm:" column:id;AUTO_INCREMENT;comment:id;not null"` // id
	KnowledgeId int       `json:"knowledge_id" gorm:"not null"`                             // 知识点id
	TypeId     int       `json:"type_id" gorm:"not null"`                                 // 题目id
	CreateAt   time.Time `json:"create_at"`                                               // 创建时间
	DeleteAt   time.Time `json:"delete_at"`
	UpdateAt   time.Time `json:"update_at"`
	IsUse      int       `gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

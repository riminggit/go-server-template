package knowledgeModel

import (
	"time"
)

type Knowledge struct {
	ID          int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	Title       string    `gorm:"type:varchar(255);not null"`                             // 知识点标题
	Analysis    string    `gorm:"type:text"`                                              // 内容，内容为富文本
	ComeFrom    string    `gorm:"type:varchar(50);not null"`                              // 来源
	CreateAt    time.Time `json:"create_at"`                                              // 创建时间
	DeleteAt    time.Time `json:"delete_at"`
	UpdateAt    time.Time `json:"update_at"`
	IsUse       int       `gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

package knowledgeModel

type Knowledge struct {
	ID       int    `json:"id" gorm:"column:id;comment:id;not null"` // id
	Title    string `gorm:"type:varchar(255);not null;unique"`       // 知识点标题
	Analysis string `gorm:"type:text"`                               // 内容，内容为富文本
	ComeFrom string `gorm:"type:varchar(50);not null"`               // 来源
	CreateAt int64  `json:"create_at"`                               // 创建时间
	DeleteAt int64  `json:"delete_at"`
	UpdateAt int64  `json:"update_at"`
	IsUse    int    `gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

package typeModel

type Type struct {
	ID         int    `json:"id" gorm:"type:bigint;column:id;comment:id;not null"` // id
	TypeName   string `json:"type_name" gorm:"type:varchar(255);not null;unique"`  // 标签名
	ClassifyId int    `json:"classify_id" gorm:"type:bigint;not null"`             // 类型id
	Rank       int    `json:"rank" gorm:"type:int(11)"`                            // 排序
	CreateAt   int64  `json:"create_at"`                                           // 创建时间
	DeleteAt   int64  `json:"delete_at"`
	UpdateAt   int64  `json:"update_at"`
	IsUse      int    `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

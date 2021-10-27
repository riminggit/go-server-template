package typeModel

type Type struct {
	ID         int    `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	TypeName   string `gorm:"type:varchar(255);not null"`                             // 标签名
	ClassifyId int    `gorm:"type:int(11);not null"`                                  // 类型id
	IsUse      int    `gorm:"type:int(3);not null;default:1"`                         // 是否删除:0 删除 1未删除
}

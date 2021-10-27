package tagModel

type Tag struct {
	ID      int    `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	TagName string `gorm:"type:varchar(50);not null"`                              // 标签名
	IsUse   int    `gorm:"type:int(3);not null;default:1"`                         // 是否删除:0 删除 1未删除
}

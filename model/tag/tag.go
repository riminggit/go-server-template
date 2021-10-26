package tagModel

type Tag struct {
	ID      int    `gorm:"autoIncrement;primaryKey;type:int(11);uniqueIndex;not null"` // id
	TagName string `gorm:"type:varchar(50);not null"`                                  // 标签名
	IsUse   int    `gorm:"type:int(3);not null"`                                       // 是否删除:0 删除 1未删除
}

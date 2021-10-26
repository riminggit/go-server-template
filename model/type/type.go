package typeModel

type Type struct {
	ID         int    `gorm:"autoIncrement;primaryKey;type:int(11);uniqueIndex;not null"` // id
	TypeName   string `gorm:"type:varchar(255);not null"`                                 // 标签名
	ClassifyId int    `gorm:"type:int(11);not null"`                                      // 类型id
	IsUse      int    `gorm:"type:int(3);not null"`                                       // 是否删除:0 删除 1未删除
}

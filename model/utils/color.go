package utils

type Color struct {
	ID    int    `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	Color string `json:"color" gorm:"type:varchar(14);not null"`                 // 颜色
	IsUse int    `json:"is_use" gorm:"type:int(3);not null;default:1"`           // 是否删除:0 删除 1未删除
}

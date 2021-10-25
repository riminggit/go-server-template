package classifyModel

import (
// "time"
)

type Classify struct {
	ID           int    `gorm:"autoIncrement;primaryKey;type:int(11);uniqueIndex;not null"` // id
	ClassifyName string `gorm:"type:varchar(30);not null"`                                  // 类型名：html5、css、js等
	ImgUrl       string `gorm:"type:varchar(255)"`                                          // 分类素材URL
	ImgSvg       string `gorm:"type:text"`                                                  // 素材svg
	IsUse        int    `gorm:"type:int(3);not null"`                                       // 是否删除:0 删除 1未删除
}

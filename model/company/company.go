package companyModel

import (
// "time"
)

type Company struct {
	ID          int    `gorm:"autoIncrement;primaryKey;type:int(11);uniqueIndex;not null"` // id
	CompanyName string `gorm:"type:varchar(50);not null"`                                  // 公司名
	ImgUrl      string `gorm:"type:varchar(255)"`                                          // 素材URL
	ImgSvg      string `gorm:"type:text"`                                                  // 素材svg
	IsUse       int    `gorm:"type:int(3);not null"`                                       // 是否删除:0 删除 1未删除
}

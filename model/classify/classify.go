package classifyModel

type Classify struct {
	ID           int    `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	ClassifyName string `gorm:"type:varchar(30);not null"`                              // 类型名：html5、css、js等
	ImgUrl       string `gorm:"type:varchar(255)"`                                      // 分类素材URL
	ImgSvg       string `gorm:"type:text"`                                              // 素材svg
	IsUse        int    `gorm:"type:int(3);not null;default:1"`                         // 是否删除:0 删除 1未删除
}

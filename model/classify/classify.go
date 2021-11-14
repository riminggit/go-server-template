package classifyModel

type Classify struct {
	ID           int    `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	ClassifyName string `json:"classify_name" gorm:"type:varchar(30);not null"`                              // 类型名：html5、css、js等
	ImgUrl       string `json:"img_url" gorm:"type:varchar(255)"`                                      // 分类素材URL
	ImgSvg       string `json:"img_svg" gorm:"type:text"`                                              // 素材svg
	Rank         int    `json:"rank" gorm:"type:int(11)"`                                           // 排序
	IsUse        int    `json:"is_use" gorm:"type:int(3);not null;default:1"`                         // 是否删除:0 删除 1未删除
}

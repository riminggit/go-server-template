package classifyModel

import "time"

type Classify struct {
	ID           int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	ClassifyName string    `json:"classify_name" gorm:"type:varchar(30);not null;unique"`  // 类型名：html5、css、js等
	ImgUrl       string    `json:"img_url" gorm:"type:varchar(255)"`                       // 分类素材URL
	ImgSvg       string    `json:"img_svg" gorm:"type:text"`                               // 素材svg
	Rank         int       `json:"rank" gorm:"type:int(11)"`                               // 排序
	CreateAt     time.Time `json:"create_at"`                                              // 创建时间
	DeleteAt     time.Time `json:"delete_at"`
	UpdateAt     time.Time `json:"update_at"`
	IsUse        int       `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

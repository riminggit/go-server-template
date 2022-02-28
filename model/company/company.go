package companyModel

type Company struct {
	ID          int    `json:"id" gorm:"type:bigint;column:id;comment:id;not null"`  // id
	CompanyName string `json:"company_name" gorm:"type:varchar(50);not null;unique"` // 公司名
	ImgUrl      string `json:"img_url" gorm:"type:varchar(255)"`                     // 素材URL
	ImgSvg      string `json:"img_svg" gorm:"type:text"`
	CreateAt    int64  `json:"create_at"` // 创建时间
	DeleteAt    int64  `json:"delete_at"`
	UpdateAt    int64  `json:"update_at"`                                    // 素材svg
	IsUse       int    `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

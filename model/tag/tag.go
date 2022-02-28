package tagModel

type Tag struct {
	ID       int    `json:"id" gorm:"type:bigint;column:id;comment:id;not null"` // id
	TagName  string `gorm:"type:varchar(50);not null;unique"`                    // 标签名
	CreateAt int64  `json:"create_at"`                                           // 创建时间
	DeleteAt int64  `json:"delete_at"`
	UpdateAt int64  `json:"update_at"`
	IsUse    int    `gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

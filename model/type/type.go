package typeModel

import "time"

type Type struct {
	ID         int       `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	TypeName   string    `json:"type_name" gorm:"type:varchar(255);not null;unique"`                      // 标签名
	ClassifyId int       `json:"classify_id" gorm:"type:int(11);not null"`                                  // 类型id
	Rank       int       `json:"rank" gorm:"type:int(11)"`                               // 排序
	CreateAt   time.Time `json:"create_at"`                                              // 创建时间
	DeleteAt   time.Time `json:"delete_at"`
	UpdateAt   time.Time `json:"update_at"`
	IsUse      int       `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

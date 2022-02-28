package userModel

type UserMistakeRecord struct {
	ID         int   `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	UserId     int   `gorm:"type:bigint;not null"`                                   // 用户id
	TopicId    int   `gorm:"type:bigint;not null"`                                   // 对应题目id
	MistakeNum int   `gorm:"type:int(3);not null"`                                   // 错误次数
	RightNum   int   `gorm:"type:int(3);not null"`                                   // 正确次数
	CreateAt   int64 `json:"create_at"`                                              // 创建时间
	DeleteAt   int64 `json:"delete_at"`
	UpdateAt   int64 `json:"update_at"`
	IsUse      int   `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

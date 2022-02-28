package userModel

type UserExperience struct {
	ID         int   `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	UserId     int   `json:"user_id" gorm:"type:bigint;not null"`                    // 用户id
	Experience int   `json:"experience" gorm:"type:bigint;not null"`                 // 用户经验
	Level      int   `json:"level" gorm:"type:int(3);not null"`                      // 用户等级
	CanTest    int   `json:"can_test" gorm:"type:int(3);not null"`                   // 是否可以参加考试  0 否 1 是，参加考试后变回0
	CreateAt   int64 `json:"create_at"`                                              // 创建时间
	DeleteAt   int64 `json:"delete_at"`
	UpdateAt   int64 `json:"update_at"`
	IsUse      int   `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

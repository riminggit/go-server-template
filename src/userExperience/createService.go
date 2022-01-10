package userExperience

import (
	"go-server-template/model/user"
	"gorm.io/gorm"
	"time"
)

// 用户新建账户的时候新建
func CreateService(userId int, db *gorm.DB) error {
	var delErr error

	createData := &userModel.UserExperience{
		UserId:     userId,
		Experience: 0,
		Level:      1,
		CreateAt:   time.Now().Add(8 * time.Hour),
		IsUse:      1,
	}

	delErr = db.Model(&userModel.UserExperience{}).Create(createData).Error
	return delErr
}

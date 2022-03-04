package userExperience

import (
	userModel "go-server-template/model/user"
	DB "go-server-template/pkg/db"
	util "go-server-template/pkg/utils"

	"gorm.io/gorm"
)

// 用户新建账户的时候新建
func CreateService(userId int, db *gorm.DB) error {
	var delErr error

	createData := &userModel.UserExperience{
		UserId:     userId,
		Experience: 0,
		Level:      1,
		CreateAt:   util.GetNowTimeUnix(),
		IsUse:      1,
	}

	delErr = db.Model(&userModel.UserExperience{}).Create(createData).Error
	return delErr
}

func CreateServiceNoAffair(userId int) error {
	var delErr error

	createData := &userModel.UserExperience{
		UserId:     userId,
		Experience: 0,
		Level:      1,
		CreateAt:   util.GetNowTimeUnix(),
		IsUse:      1,
	}

	delErr = DB.DBLivingExample.Model(&userModel.UserExperience{}).Create(createData).Error
	return delErr
}

package userExperience


import (
	"go-server-template/model/user"
	"gorm.io/gorm"

)

func DeleteService(userId int, db *gorm.DB) error {
	return db.Delete(&userModel.UserExperience{},"user_id = ?", userId).Error
}

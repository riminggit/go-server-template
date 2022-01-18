package userExperience

import (
	"github.com/gin-gonic/gin"
	"go-server-template/model/user"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/utils"
	"gorm.io/gorm"
	"time"
)


// 用户新增经验
func UserAddExExperience(params UpdateParams, c *gin.Context, tx *gorm.DB) bool {
	res := true

	userId := util.GetUserInfo(c).Data.ID
	experienceAdd := Exalculate(params)

	queryInfo := UserQueryExperienceSimple(c, userId).Data

	canAddEx := JudgeUserCanAddExperience(&queryInfo, experienceAdd)

	setData := userModel.UserExperience{
		UpdateAt: time.Now().Add(8 * time.Hour),
	}

	if canAddEx.CanAddEx {
		setData.Experience = experienceAdd
	} else {
		setData.CanTest = 1
		setData.Experience = canAddEx.LevelMaxEx
	}

	up := DB.DBLivingExample.Model(&userModel.UserExperience{})
	up = up.Where("user_id = ?", userId)
	up = up.Updates(setData)
	if up.Error != nil {
		res = false
	}

	return res
}


// 用户考试成功后修改经验数据
func UserTestSuccessCanAddExperience(c *gin.Context,tx *gorm.DB,userId int) error {
	var Err error

	queryInfo := UserQueryExperienceSimple(c, userId).Data
	canAddEx := JudgeUserCanAddExperience(&queryInfo, 0)
	setData := userModel.UserExperience{
	  UpdateAt: time.Now().Add(8 * time.Hour),
	}
	if !canAddEx.CanAddEx {
		setData.Level = queryInfo.Level + 1
		setData.CanTest = 0
	} 

	up := DB.DBLivingExample.Model(&userModel.UserExperience{})
	up = up.Where("user_id = ?", userId)
	up = up.Updates(setData)
	if up.Error != nil {
		tx.Rollback()
		Err = up.Error
	}

	return Err
}

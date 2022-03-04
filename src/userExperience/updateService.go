package userExperience

import (

	userModel "go-server-template/model/user"
	DB "go-server-template/pkg/db"
	util "go-server-template/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 用户新增经验
func UserAddExExperience(params UpdateParams, c *gin.Context) bool {
	res := true

	userEX := &userModel.UserExperience{}

	up := DB.DBLivingExample.Model(&userModel.UserExperience{})
	up = up.Where("user_id = ?", params.UserId).Find(userEX)

	if userEX.ID == 0 {
		CreateServiceNoAffair(params.UserId)
	} else {
		experienceAdd := Exalculate(params)

		queryInfo := UserQueryExperienceSimple(c, params.UserId).Data

		canAddEx := JudgeUserCanAddExperience(&queryInfo, experienceAdd)

		setData := &userModel.UserExperience{}

		if canAddEx.CanAddEx {
			setData.Experience = userEX.Experience + experienceAdd
		} else {
			setData.CanTest = 1
			setData.Experience = canAddEx.LevelMaxEx
		}
		setData.UpdateAt = util.GetNowTimeUnix()
		up = up.Updates(setData)
	}

	if up.Error != nil {
		res = false
	}

	return res
}

// 用户答套题新增经验
func UserAddExExperienceTopicSet(params UpdateTopicSetParams, c *gin.Context) bool {
	res := true

	userEX := &userModel.UserExperience{}

	up := DB.DBLivingExample.Model(&userModel.UserExperience{})
	up = up.Where("user_id = ?", params.UserId).Find(userEX)

	if userEX.ID == 0 {
		CreateServiceNoAffair(params.UserId)
	} else {
		experienceAdd := ExalculateTopicSet(params)

		queryInfo := UserQueryExperienceSimple(c, params.UserId).Data
	
		canAddEx := JudgeUserCanAddExperience(&queryInfo, experienceAdd)

		setData := &userModel.UserExperience{}

		if canAddEx.CanAddEx {
			setData.Experience = userEX.Experience + experienceAdd
		} else {
			setData.CanTest = 1
			setData.Experience = canAddEx.LevelMaxEx
		}

		setData.UpdateAt = util.GetNowTimeUnix()
		up = up.Updates(setData)
	}

	if up.Error != nil {
		res = false
	}

	return res
}

// 用户考试成功后修改经验数据
func UserTestSuccessCanAddExperience(c *gin.Context, tx *gorm.DB, userId int) error {
	var Err error

	queryInfo := UserQueryExperienceSimple(c, userId).Data
	canAddEx := JudgeUserCanAddExperience(&queryInfo, 0)
	setData := userModel.UserExperience{
		UpdateAt: util.GetNowTimeUnix(),
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

package feedback

import (
	"github.com/gin-gonic/gin"
	"go-server-template/model/user"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/pkg/utils"
	"time"
)

func AnswerFeedbackService(c *gin.Context, params AnswerFeedbackParams) *CommonReturn {
	res := &CommonReturn{}

	thisTime := time.Now().Add(8 * time.Hour)
	if params.ID != 0 {
		setData := userModel.UserFeedback{IsUse: 1}

		if params.FacebackAnswer != "" {
			setData.FacebackAnswer = params.FacebackAnswer
			setData.FacebackAnswerTime = thisTime
		}
		if params.IsRead != 0 {
			setData.IsRead = params.IsRead
		}

		Res := DB.DBLivingExample.Model(&userModel.UserFeedback{}).Where("id = ?", params.ID).Updates(setData)
		if Res.Error != nil {
			res.Code = e.UPDATE_DATA_FALE
		}
	} else {
		res.Code = e.NO_DATA_EXISTS
	}
	CleanRedisQuery(c)
	res.Code = e.SUCCESS
	return res
}


func UserUpdateService(c *gin.Context, params UserUpdateParams) *CommonReturn {
	res := &CommonReturn{}
	thisTime := time.Now().Add(8 * time.Hour)

	userInfoRes := util.GetUserInfo(c)
	if userInfoRes.Code != e.SUCCESS {
		res.Code = userInfoRes.Code
		return res
	}

	queryInfo := &userModel.UserFeedback{}

	if params.ID != 0 {
		setData := userModel.UserFeedback{
			IsUse:    1,
			UpdateAt: thisTime,
		}

		if params.Content != "" {
			setData.Content = params.Content
		}

		if params.Grade != 0 {
			setData.IsRead = params.Grade
		}
		dbfunc := DB.DBLivingExample.Model(&userModel.UserFeedback{}).Where("id = ?", params.ID).Find(&queryInfo)
		if queryInfo.UserId != userInfoRes.Data.ID {
			res.Code = e.NOT_EDUIT_AUTH
			return res
		}

		Res := dbfunc.Updates(setData)
		if Res.Error != nil {
			res.Code = e.UPDATE_DATA_FALE
		}
	} else {
		res.Code = e.NO_DATA_EXISTS
	}
	
	CleanRedisQuery(c)
	res.Code = e.SUCCESS
	return res
}

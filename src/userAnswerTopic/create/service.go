package userAnswerTopicCreate

import (
	userModel "go-server-template/model/user"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"
	topicSetQuery "go-server-template/src/topicSet/query"

	"github.com/gin-gonic/gin"
)

func CreateService(c *gin.Context, params UserAnswerTopicCreateParams) *CreateReturn {
	res := &CreateReturn{}
	res.Code = e.SUCCESS
	res.Msg = "记录新增成功"

	if params.TopicIdList == "" && params.TopicSetId == 0 {
		res.Code = e.CREATE_DATA_FALE
		res.Msg = "题目或套题不可为空"
		return res
	}

	userInfoRes := util.GetUserInfo(c)
	if userInfoRes.Code != e.SUCCESS {
		res.Code = e.CREATE_DATA_FALE
		res.Msg = "获取用户信息失败"
		return res
	}

	DBFun := DB.DBLivingExample.Model(&userModel.UserAnswerTopicRecord{})

	userAnswerTopicRecord := &userModel.UserAnswerTopicRecord{
		UserId:      userInfoRes.Data.ID,
		IsAchieve:   0,
		CreateAt:    util.GetNowTimeUnix(),
		IsUse:       1,
		AnswerStart: util.GetNowTimeUnix(),
	}

	if params.TopicSetId > 0 {
		userAnswerTopicRecord.TopicSetId = params.TopicSetId
		queryTopicSetParams := &topicSetQuery.QueryTopicSetSingleParams{
			ID: params.TopicSetId,
		}
		topicSet := topicSetQuery.QueryTopicSetSimpleSingleService(c, *queryTopicSetParams)
		if topicSet.Code == e.SUCCESS {
			resData := topicSet.Data
			userAnswerTopicRecord.AnswerNum = resData.TopicNum
			userAnswerTopicRecord.TopicDifficulty = resData.TopicSetDifficulty
			userAnswerTopicRecord.TopicLevel = resData.TopicSetLevel
		}
	}

	if params.TopicIdList != "" {
		userAnswerTopicRecord.TopicIdList = params.TopicIdList
		if params.AnswerNum > 0 {
			userAnswerTopicRecord.AnswerNum = params.AnswerNum
		}

		if params.TopicDifficulty > 0 {
			userAnswerTopicRecord.TopicDifficulty = params.TopicDifficulty
		}

		if params.TopicLevel > 0 {
			userAnswerTopicRecord.TopicLevel = params.TopicLevel
		}
	}

	if params.Remark != "" {
		userAnswerTopicRecord.Remark = params.Remark
	}

	// 查询数据，存在则更新，不存在则新增
	queryRecord := &userModel.UserAnswerTopicRecord{}
	DBFun.Where("user_id = ?", userInfoRes.Data.ID).Where("topic_set_id = ?", params.TopicSetId).Find(&queryRecord)

	var err error
	if queryRecord.ID > 0 {
		err = DBFun.Updates(userAnswerTopicRecord).Error
	} else {
		err = DBFun.Create(userAnswerTopicRecord).Error
	}

	if err != nil {
		res.Code = e.CREATE_DATA_FALE
		res.Msg = err.Error()
	}

	return res
}

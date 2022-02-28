package userAnswerTopicUpdate

import (
	userModel "go-server-template/model/user"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"

	"github.com/gin-gonic/gin"
)

func UpdateService(c *gin.Context, params UserAnswerTopicUpdateParams) *UpdateReturn {
	res := &UpdateReturn{}
	res.Code = e.SUCCESS

	userInfoRes := util.GetUserInfo(c)
	if userInfoRes.Code != e.SUCCESS {
		res.Code = e.CREATE_DATA_FALE
		res.Msg = "记录新增成功"
		return res
	}

	userAnswerTopicRecord := &userModel.UserAnswerTopicRecord{
		IsAchieve: 1,
		UpdateAt:   util.GetNowTimeUnix(),
	}

	// if params.TopicIdList != "" {
	// 	userAnswerTopicRecord.TopicIdList = params.TopicIdList
	// }

	// if params.TopicSetId != "" {
	// 	userAnswerTopicRecord.TopicSetId = params.TopicSetId
	// }

	// 以客户端结束时间为准
	// if !params.AnswerEnd.IsZero() {
	// 	userAnswerTopicRecord.AnswerEnd = params.AnswerEnd
	// } else {
	// 	res.Msg = "结束时间不能为空"
	// 	res.Code = e.CREATE_DATA_FALE
	// }

	if params.AnswerCorrectNum > 0 {
		userAnswerTopicRecord.AnswerNum = params.AnswerCorrectNum
	}

	if params.Remark != "" {
		userAnswerTopicRecord.Remark = params.Remark
	}

	return res
}

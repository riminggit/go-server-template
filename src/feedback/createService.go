package feedback

import (
	userModel "go-server-template/model/user"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"

	"github.com/gin-gonic/gin"
)

func CreateService(c *gin.Context, params CreateParams) *CommonReturn {
	res := &CommonReturn{}

	userInfoRes := util.GetUserInfo(c)
	if userInfoRes.Code != e.SUCCESS {
		res.Code = userInfoRes.Code
		return res
	}

	createData := &userModel.UserFeedback{
		UserId:       userInfoRes.Data.ID,
		TopicId:      params.TopicId,
		FeedbackType: params.FeedbackType,
		Content:      params.Content,
		Grade:        params.Grade,
		CreateAt:     util.GetNowTimeUnix(),
		IsUse:        1,
	}

	err := DB.DBLivingExample.Model(&userModel.UserFeedback{}).Create(createData).Error
	if err != nil {
		res.Code = e.FEEDBACK_FALE
		return res
	}

	res.Code = e.SUCCESS
	return res
}

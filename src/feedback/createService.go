package feedback

import (
	"go-server-template/model/user"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"time"
	"github.com/gin-gonic/gin"
)

func CreateService(c *gin.Context, params CreateParams) *CommonReturn {
	res := &CommonReturn{}

	createData := &userModel.UserFeedback{
		CreateAt: time.Now().Add(8 * time.Hour),
		IsUse: 1,
		TopicId:params.TopicId,
		FeedbackType:params.FeedbackType,
		FacebackAnswer:params.FacebackAnswer,
		Content:params.Content,
		Grade:params.Grade,
	}

	err := DB.DBLivingExample.Model(&userModel.UserFeedback{}).Create(createData).Error
	if err != nil {
		res.Code = e.FEEDBACK_FALE
		return res
	}

	res.Code = e.SUCCESS
	return res
}

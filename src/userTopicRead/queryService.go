package userTopicRead

import (
	userModel "go-server-template/model/user"
	DB "go-server-template/pkg/db"

	"github.com/gin-gonic/gin"
)

func QueryUserTopicReadSimple(params *QuerSimpleParams, c *gin.Context) *userModel.UserTopicRead {
	data := &userModel.UserTopicRead{}

	queryFun := DB.DBLivingExample
	queryFun = queryFun.Where("user_id = ?", params.UserId).Where("topic_id = ?", params.TopicId)
	queryFun.Model(&userModel.UserExperience{}).First(&data)

	return data
}

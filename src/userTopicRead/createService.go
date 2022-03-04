package userTopicRead

import (
	"github.com/gin-gonic/gin"

	userModel "go-server-template/model/user"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"
	topicQuery "go-server-template/src/topic/query"
	"go-server-template/src/userExperience"
)

func CreateUserTopicRead(c *gin.Context, params CreateParams) *CreateReturn {
	res := &CreateReturn{}
	res.Code = e.SUCCESS
	res.Msg = "记录新增成功"

	userInfoRes := util.GetUserInfo(c)
	if userInfoRes.Code != e.SUCCESS {
		res.Code = e.CREATE_DATA_FALE
		res.Msg = "获取用户信息失败"
		return res
	}

	userRecord := &userModel.UserTopicRead{}

	CreateFun := DB.DBLivingExample.Model(&userModel.UserTopicRead{})
	CreateFun = CreateFun.Where("user_id = ?", userInfoRes.Data.ID).Where("topic_id = ?", params.TopicId)
	CreateFun.Find(userRecord)

	if userRecord.ID == 0 {
		createParams := &userModel.UserTopicRead{
			UserId:   userInfoRes.Data.ID,
			TopicId:  params.TopicId,
			IsRead:   1,
			ReadNum:  1,
			CreateAt: util.GetNowTimeUnix(),
		}
		CErr := CreateFun.Model(&userModel.UserTopicRead{}).Create(&createParams).Error
		if CErr != nil {
			res.Code = e.CREATE_DATA_FALE
			res.Msg = "记录新增失败"
			return res
		}
	} else {

		updateParams := &userModel.UserTopicRead{
			IsRead:   1,
			ReadNum:  userRecord.ReadNum + 1,
			UpdateAt: util.GetNowTimeUnix(),
		}
		CreateFun.Model(&userModel.UserTopicRead{}).Updates(&updateParams)
	}

	queryTopicParams := &topicQuery.QueryTopicSimpleSingleParams{
		Id: params.TopicId,
	}

	topicInfo := topicQuery.QueryTopicSimpleSingleService(c, *queryTopicParams).Data

	addExParams := &userExperience.UpdateParams{
		UserId:           userInfoRes.Data.ID,
		Degree:           topicInfo.Degree,
		Level:            topicInfo.Level,
		IsBaseTopic:      topicInfo.IsBaseTopic,
		IsImportantTopic: topicInfo.IsImportantTopic,
	}

	ExRes := userExperience.UserAddExExperience(*addExParams, c)

	if !ExRes {
		res.Code = e.CREATE_DATA_FALE
		res.Msg = "记录新增失败"
	}

	return res
}

package userTopicRead

import (
	"time"

	"github.com/gin-gonic/gin"

	// DB "go-server-template/pkg/db"
	userModel "go-server-template/model/user"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"

	"gorm.io/gorm"
)

func CreateUserTopicRead(params CreateParams, c *gin.Context, tx *gorm.DB) error {
	var Err error

	userInfoRes := util.GetUserInfo(c)
	if userInfoRes.Code != e.SUCCESS {
		return Err
	}

	userRecord := &userModel.UserTopicRead{}

	CreateErr := tx.Model(&userModel.UserExperience{})
	CreateErr = CreateErr.Where("user_id = ?", userInfoRes.Data.ID).Where("topic_id = ?", params.TopicId)
	CreateErr = CreateErr.First(userRecord)

	if userRecord.ID == 0 {
		createParams := userModel.UserTopicRead{
			UserId:   userInfoRes.Data.ID,
			TopicId:  params.TopicId,
			IsRead:   1,
			ReadNum:  1,
			CreateAt: time.Now().Add(8 * time.Hour),
		}
		CreateErr.Create(createParams)
	} else {
		updateParams := userModel.UserTopicRead{
			IsRead:   1,
			ReadNum:  userRecord.ReadNum + 1,
			UpdateAt: time.Now().Add(8 * time.Hour),
		}
		CreateErr.Updates(updateParams)
	}
	return Err
}

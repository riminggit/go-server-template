package userTopicRead

import (
	userModel "go-server-template/model/user"
	DB "go-server-template/pkg/db"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUserTopicReadSimple(params *UpdateSimpleParams, c *gin.Context, tx *gorm.DB) {
	updataParams := userModel.UserTopicRead{
		IsRead:   1,
		ReadNum:  params.ReadNum + 1,
		UpdateAt: time.Now().Add(8 * time.Hour),
	}

	up := DB.DBLivingExample.Model(&userModel.UserExperience{})
	up = up.Where("user_id = ?", params.UserId).Where("topic_id = ?", params.TopicId)
	up = up.Updates(updataParams)
	if up.Error != nil {
		tx.Rollback()
	}
}

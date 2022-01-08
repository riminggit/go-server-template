package feedback

import (
	"github.com/gin-gonic/gin"
	"go-server-template/model/user"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	"time"
)

func DeleteService(c *gin.Context, params DeleteParams) *CommonReturn {
	res := &CommonReturn{}

	if params.RealDel == "1" {
		if len(params.ID) > 0 {
			delErr := DB.DBLivingExample.Delete(&userModel.UserFeedback{}, "id IN ?", params.ID).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Code = e.DELETE_DATA_FALE
				return res
			}
		}
		if len(params.UserId) > 0 {
			delErr := DB.DBLivingExample.Delete(&userModel.UserFeedback{}, "user_id IN ?", params.UserId).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Code = e.DELETE_DATA_FALE
				return res
			}
		}
		if len(params.TopicId) > 0 {
			delErr := DB.DBLivingExample.Delete(&userModel.UserFeedback{}, "topic_id IN ?", params.TopicId).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Code = e.DELETE_DATA_FALE
				return res
			}
		}
		if len(params.CreateAt) > 0 {
			delErr := DB.DBLivingExample.Delete(&userModel.UserFeedback{}, "create_at between ? and ?", params.CreateAt[0], params.CreateAt[1]).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Code = e.DELETE_DATA_FALE
				return res
			}
		}
		if len(params.DeleteAt) > 0 {
			delErr := DB.DBLivingExample.Delete(&userModel.UserFeedback{}, "delete_at between ? and ?", params.CreateAt[0], params.CreateAt[1]).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Code = e.DELETE_DATA_FALE
				return res
			}
		}
		if len(params.UpdateAt) > 0 {
			delErr := DB.DBLivingExample.Delete(&userModel.UserFeedback{}, "update_at between ? and ?", params.CreateAt[0], params.CreateAt[1]).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Code = e.DELETE_DATA_FALE
				return res
			}
		}
	} else {
		setData := map[string]interface{}{"is_use": 0, "DeleteAt": time.Now().Add(8 * time.Hour)}
		if len(params.ID) > 0 {
			delErr := DB.DBLivingExample.Model(&userModel.UserFeedback{}).Where("id IN ?", params.ID).Updates(setData).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Code = e.DELETE_DATA_FALE
				return res
			}
		}
		if len(params.UserId) > 0 {
			delErr := DB.DBLivingExample.Model(&userModel.UserFeedback{}).Where("user_id IN ?", params.ID).Updates(setData).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Code = e.DELETE_DATA_FALE
				return res
			}
		}
		if len(params.TopicId) > 0 {
			delErr := DB.DBLivingExample.Model(&userModel.UserFeedback{}).Where("topic_id IN ?", params.ID).Updates(setData).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Code = e.DELETE_DATA_FALE
				return res
			}
		}
		if len(params.CreateAt) > 0 {
			delErr := DB.DBLivingExample.Model(&userModel.UserFeedback{}).Where("create_at between ? and ?", params.CreateAt[0], params.CreateAt[1]).Updates(setData).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Code = e.DELETE_DATA_FALE
				return res
			}
		}
		if len(params.DeleteAt) > 0 {
			delErr := DB.DBLivingExample.Model(&userModel.UserFeedback{}).Where("delete_at between ? and ?", params.CreateAt[0], params.CreateAt[1]).Updates(setData).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Code = e.DELETE_DATA_FALE
				return res
			}
		}
		if len(params.UpdateAt) > 0 {
			delErr := DB.DBLivingExample.Model(&userModel.UserFeedback{}).Where("update_at between ? and ?", params.CreateAt[0], params.CreateAt[1]).Updates(setData).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Code = e.DELETE_DATA_FALE
				return res
			}
		}
	}

	res.Code = e.SUCCESS
	return res
}

func UserDeleteService(c *gin.Context, params UserDeleteParams) *CommonReturn {
	res := &CommonReturn{}

	setData := map[string]interface{}{"is_use": 0, "DeleteAt": time.Now().Add(8 * time.Hour)}
	if len(params.ID) > 0 {
		delErr := DB.DBLivingExample.Model(&userModel.UserFeedback{}).Where("id IN ?", params.ID).Updates(setData).Error
		if delErr != nil {
			logging.Error(delErr)
			res.Code = e.DELETE_DATA_FALE
			return res
		}
	}

	res.Code = e.SUCCESS
	return res
}

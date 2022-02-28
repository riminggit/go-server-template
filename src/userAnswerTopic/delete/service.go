package userAnswerTopicDelete

import (
	userModel "go-server-template/model/user"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"

	"github.com/gin-gonic/gin"
)

func DeleteService(c *gin.Context, params DeleteParams) *DeleteReturn {
	res := &DeleteReturn{}
	res.Code = e.SUCCESS
	res.Msg = "操作成功"

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
	DBFun = DBFun.Where("user_id = ?", userInfoRes.Data.ID)

	if params.TopicSetId > 0 {
		DBFun = DBFun.Where("topic_set_id = ?", params.TopicSetId)
	} else {
		DBFun = DBFun.Where("topic_id_list = ?", params.TopicIdList)
	}

	err := DBFun.Update("is_use", 0).Update("delete_at", util.GetNowTimeUnix()).Error

	if err != nil {
		res.Code = e.CREATE_DATA_FALE
		res.Msg = err.Error()
	}

	return res
}

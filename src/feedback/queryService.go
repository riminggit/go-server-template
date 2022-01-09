package feedback

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-server-template/config"
	"go-server-template/model/user"
	"go-server-template/pkg/apiMap"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
	"go-server-template/pkg/utils"
)

func QueryService(c *gin.Context, params QueryParams) *queryPaddingReturn {
	res := &queryPaddingReturn{}
	var queryInfo []userModel.UserFeedback
	var RGetData queryPaddingReturnData

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_FEEDBACK)
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &queryInfo)
		if err != nil {
			logging.Debug(err)
		}
		res.Code = e.SUCCESS
		res.Data = RGetData
		return res
	}

	queryFun := DB.DBLivingExample.Where("is_use = ?", params.IsUse)
	if len(params.Id) > 0 {
		queryFun = queryFun.Where("id IN (?)", params.Id)
	}

	if len(params.UserId) > 0 {
		queryFun = queryFun.Where("user_id IN ?", params.UserId)
	}

	if params.TopicId != 0 {
		queryFun = queryFun.Where("topic_id = ?", params.TopicId)
	}

	if params.FeedbackType != 0 {
		queryFun = queryFun.Where("feedback_type = ?", params.FeedbackType)
	}

	if params.Grade != 0 {
		queryFun = queryFun.Where("grade = ?", params.Grade)
	}

	if len(params.CreateAt) > 0 {
		queryFun = queryFun.Where("create_at between ? and ?", params.CreateAt[0], params.CreateAt[1])
	}

	if len(params.DeleteAt) > 0 {
		queryFun = queryFun.Where("delete_at between ? and ?", params.DeleteAt[0], params.DeleteAt[1])
	}

	if len(params.UpdateAt) > 0 {
		queryFun = queryFun.Where("update_at between ? and ?", params.UpdateAt[0], params.UpdateAt[1])
	}

	queryFun = queryFun.Limit(params.PageSize).Offset((params.PageNum - 1) * params.PageSize)
	Err := queryFun.Model(&userModel.UserFeedback{}).Find(&queryInfo).Count(&res.Data.PagingArgument.Total).Error
	if Err != nil {
		res.Code = e.NO_DATA_EXISTS
		return res
	}

	res.Data.PagingArgument.PageNum = params.PageNum
	res.Data.PagingArgument.PageSize = params.PageSize
	res.Data.Data = queryInfo

	RSetData := res.Data
	Redis.SetValue(queryRedisParams, RSetData, dataRxpirationTime)

	res.Code = e.SUCCESS

	return res
}

func UserQueryService(c *gin.Context, params UserQueryParams) *CommonQueryReturn {
	res := &CommonQueryReturn{}
	var queryInfo []userModel.UserFeedback

	getUserInfo := util.GetUserPrefix(c)
	userPrefix := getUserInfo.UserPrefix
	userInfo := getUserInfo.UserData
	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_FEEDBACK) + userPrefix
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &queryInfo)
		if err != nil {
			logging.Debug(err)
		}
		res.Code = e.SUCCESS
		res.Data = queryInfo
		return res
	}

	queryFun := DB.DBLivingExample.Where("is_use = ?", params.IsUse)
	queryFun = queryFun.Where("user_id = ?", userInfo.ID)

	if len(params.Id) > 0 {
		queryFun = queryFun.Where("id IN (?)", params.Id)
	}

	if params.TopicId != 0 {
		queryFun = queryFun.Where("topic_id = ?", params.TopicId)
	}

	if params.FeedbackType != 0 {
		queryFun = queryFun.Where("feedback_type = ?", params.FeedbackType)
	}

	if params.Grade != 0 {
		queryFun = queryFun.Where("grade = ?", params.Grade)
	}

	if len(params.CreateAt) > 0 {
		queryFun = queryFun.Where("create_at between ? and ?", params.CreateAt[0], params.CreateAt[1])
	}

	if len(params.DeleteAt) > 0 {
		queryFun = queryFun.Where("delete_at between ? and ?", params.DeleteAt[0], params.DeleteAt[1])
	}

	if len(params.UpdateAt) > 0 {
		queryFun = queryFun.Where("update_at between ? and ?", params.UpdateAt[0], params.UpdateAt[1])
	}

	Err := queryFun.Model(&userModel.UserFeedback{}).Find(&queryInfo).Error
	if Err != nil {
		res.Code = e.NO_DATA_EXISTS
		return res
	}

	res.Data = queryInfo

	Redis.SetValue(queryRedisParams, queryInfo, dataRxpirationTime)

	res.Code = e.SUCCESS

	return res
}

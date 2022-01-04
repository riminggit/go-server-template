package midTopicClassifyQuery

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-server-template/config"
	"go-server-template/model/topic"
	"go-server-template/pkg/apiMap"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
)

func QueryTopicClassifyMid(c *gin.Context, params QueryTopicClassifyMidParams) *QueryTopicClassifyMidReturn {

	res := &QueryTopicClassifyMidReturn{}
	var queryInfo []topicModel.TopicClassify

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC_CLASSIFY_MID)
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



	queryFun := DB.DBLivingExample
	if len(params.Id) > 0 {
		queryFun = queryFun.Where("id IN (?)", params.Id)
	}

	if params.TopicId != "" {
		queryFun = queryFun.Where("topic_id = ?", params.TopicId)
	}

	if params.ClassifyId != "" {
		queryFun = queryFun.Where("classify_id = ?", params.ClassifyId)
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

	queryFun.Model(&topicModel.TopicClassify{}).Find(&queryInfo)

	res.Data = queryInfo

	Redis.SetValue(queryRedisParams, queryInfo, dataRxpirationTime)

	res.Code = e.SUCCESS

	return res
}

func QueryTopicClassifyMidPading(c *gin.Context, params QueryTopicClassifyMidPadingParams) *QueryTopicClassifyMidPadingReturn {

	res := &QueryTopicClassifyMidPadingReturn{}
	var queryInfo []topicModel.TopicClassify
	var RGetData QueryTopicClassifyMidPadingReturnData

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC_CLASSIFY_MID_PADING)
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &RGetData)
		if err != nil {
			logging.Debug(err)
		}
		res.Code = e.SUCCESS
		res.Data = RGetData
		return res
	}


	if params.PageNum == 0 {
		params.PageNum = 1
	}
	if params.PageSize == 0 {
		params.PageSize = projectConfig.AppConfig.BaseConfig.PAGE_SIZE
	}

	queryFun := DB.DBLivingExample

	if len(params.TopicId) > 0 {
		queryFun = queryFun.Where("topic_id IN (?)", params.TopicId)
	}

	if len(params.ClassifyId) > 0 {
		queryFun = queryFun.Where("classify_id IN (?)", params.ClassifyId)
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
	queryFun.Model(&topicModel.TopicClassify{}).Find(&queryInfo).Count(&res.Data.PagingArgument.Total)

	res.Data.PagingArgument.PageNum = params.PageNum
	res.Data.PagingArgument.PageSize = params.PageSize
	res.Data.Data = queryInfo

	RSetData := res.Data

	Redis.SetValue(queryRedisParams, RSetData, dataRxpirationTime)

	res.Code = e.SUCCESS

	return res
}

package topicQuery

import (
	"encoding/json"
	"go-server-template/config"
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
	"github.com/gin-gonic/gin"
	// "go-server-template/src/classify/query"
	// "go-server-template/model/classify"
)


func queryTopicService(c *gin.Context, params queryTopicParams) *queryTopicReturn {
	res := &queryTopicReturn{}
	var queryInfo []topicModel.Topic
	var RGetData TopicReturnData

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := "query-topic:"
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

	queryFun := DB.DBLivingExample.Where("is_use = ?", params.IsUse)
	if len(params.Id) > 0 {
		queryFun = queryFun.Where("id IN (?)", params.Id)
	}

	if params.Title != "" {
		queryFun = queryFun.Where("title = ?", params.Title)
	}

	if len(params.QuestionType) > 0 {
		queryFun = queryFun.Where("question_type IN (?)", params.QuestionType)
	}

	if len(params.Degree) > 0 {
		queryFun = queryFun.Where("degree IN (?)", params.Degree)
	}

	if len(params.Level) > 0 {
		queryFun = queryFun.Where("level IN (?)", params.Level)
	}

	if params.IsBaseTopic != "" {
		queryFun = queryFun.Where("is_base_topic = ?", params.IsBaseTopic)
	}

	if params.IsImportantTopic != "" {
		queryFun = queryFun.Where("is_important_topic = ?", params.IsImportantTopic)
	}

	if len(params.ComeFrom) > 0 {
		queryFun = queryFun.Where("come_from IN (?)", params.ComeFrom)
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

	queryFun.Model(&topicModel.Topic{}).Find(&queryInfo).Count(&res.Data.PagingArgument.Total)

	// for  _, item := range queryInfo {
	// 	queryClassifyData := classifyQuery.QueryClassifyService(c, rParams)
	// }

	res.Data.PagingArgument.PageNum = params.PageNum
	res.Data.PagingArgument.PageSize = params.PageSize
	res.Data.Data = queryInfo

	RSetData := res.Data
	Redis.SetValue(queryRedisParams, RSetData, dataRxpirationTime)

	res.Code = e.SUCCESS

	return res
}



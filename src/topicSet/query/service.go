package topicSetQuery

import (
	"encoding/json"
	"go-server-template/config"
	"go-server-template/model/topic"
	"go-server-template/pkg/apiMap"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
	"github.com/gin-gonic/gin"
)

func QueryTopicSetService(c *gin.Context, params QueryTopicSetParams) *queryTopicSetReturn {
	res := &queryTopicSetReturn{}
	var queryInfo []topicModel.TopicSet

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName :=  apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC_SET)
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
	if len(params.Id) > 0 {
		queryFun = queryFun.Where("id IN (?)", params.Id)
	}

	if params.Name != "" {
		queryFun = queryFun.Where("name = ?", params.Name)
	}

	if params.TopicSetIdList != "" {
		queryFun = queryFun.Where("topic_set_id_list = ?", params.TopicSetIdList)
	}

	if params.TopicSetDifficulty != "" {
		queryFun = queryFun.Where("topic_set_difficulty = ?", params.TopicSetDifficulty)
	}

	if params.TopicSetLevel != "" {
		queryFun = queryFun.Where("topic_set_level = ?", params.TopicSetLevel)
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

	queryFun.Model(&topicModel.TopicSet{}).Find(&queryInfo)

	res.Data = queryInfo

	Redis.SetValue(queryRedisParams, queryInfo, dataRxpirationTime)

	res.Code = e.SUCCESS

	return res
}

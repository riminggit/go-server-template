package topicQuery

import (
	"encoding/json"
	"go-server-template/config"
	"go-server-template/model/topic"
	"go-server-template/pkg/apiMap"
	DB "go-server-template/pkg/db"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
)

func GetTopicCountByDegree(degree int) *int64 {

	var topicCount int64
	params := &QueryTopicCount{
		Degree: degree,
	}

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.QUERY_TOPIC_COUNT_BY_DEGREE)
	queryRedisParams := interfaceName + string(redisParamsJson)
	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &topicCount)
		if err != nil {
			logging.Debug(err)
		}
		return &topicCount
	}

	queryFun := DB.DBLivingExample.Where("is_use = ?", 1)
	queryFun = queryFun.Where("degree = ?", degree)
	queryFun.Model(&topicModel.Topic{}).Count(&topicCount)

	Redis.SetValue(queryRedisParams, topicCount, dataRxpirationTime)

	return &topicCount
}

func GetTopicCountByLevel(level int) *int64 {

	var topicCount int64
	params := &QueryTopicCount{
		Degree: level,
	}

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.QUERY_TOPIC_COUNT_BY_LEVEL)
	queryRedisParams := interfaceName + string(redisParamsJson)
	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &topicCount)
		if err != nil {
			logging.Debug(err)
		}
		return &topicCount
	}

	queryFun := DB.DBLivingExample.Where("is_use = ?", 1)
	queryFun = queryFun.Where("level = ?", level)
	queryFun.Model(&topicModel.Topic{}).Count(&topicCount)

	Redis.SetValue(queryRedisParams, topicCount, dataRxpirationTime)

	return &topicCount
}

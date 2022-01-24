package topicSetQuery

import (
	"github.com/gin-gonic/gin"
	"go-server-template/config"
	topicModel "go-server-template/model/topic"
	"go-server-template/pkg/apiMap"
	DB "go-server-template/pkg/db"
	Redis "go-server-template/pkg/redis"
	"strconv"
)

func QueryTopicSetAllCount(c *gin.Context) int64 {

	var count int64
	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME

	interfaceName := apiMap.GetRedisPrefixName(apiMap.TOPIC_SET_ALL_COUNT)
	queryRedisParams := interfaceName

	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		intData, _ := strconv.ParseInt(redisData, 8, 64)
		return intData
	}

	DB.DBLivingExample.Where("is_use = ?", 1).Model(&topicModel.TopicSet{}).Count(&count)
	Redis.SetValue(queryRedisParams, count, dataRxpirationTime)
	return count
}

func QueryTopicSetIDMin(c *gin.Context) int {

	var count int
	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME

	interfaceName := apiMap.GetRedisPrefixName(apiMap.TOPIC_SET_ID_MIN)
	queryRedisParams := interfaceName

	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		intData, _ := strconv.Atoi(redisData)
		return intData
	}

	DB.DBLivingExample.Raw("SELECT MIN(id) FROM topic_set WHERE is_use = ?", 1).Scan(&count)
	Redis.SetValue(queryRedisParams, count, dataRxpirationTime)
	return count
}

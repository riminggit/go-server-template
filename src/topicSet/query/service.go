package topicSetQuery

import (
	"encoding/json"
	projectConfig "go-server-template/config"
	topicModel "go-server-template/model/topic"
	"go-server-template/pkg/apiMap"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
	topicQuery "go-server-template/src/topic/query"
	"strings"

	"github.com/gin-gonic/gin"
)

func QueryTopicSetService(c *gin.Context, params QueryTopicSetParams) *queryTopicSetReturn {
	res := &queryTopicSetReturn{}
	var queryInfo []topicModel.TopicSet
	var queryInfoReturn []TopicInfoReturnData
	var RGetData queryTopicSetReturnData

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC_SET)
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

	if params.Name != "" {
		queryFun = queryFun.Where("name = ?", params.Name)
	}

	if params.TopicIdList != "" {
		queryFun = queryFun.Where("topic_id_list = ?", params.TopicIdList)
	}

	if params.TopicSetDifficulty != "" {
		queryFun = queryFun.Where("topic_set_difficulty = ?", params.TopicSetDifficulty)
	}

	if params.TopicSetLevel != "" {
		queryFun = queryFun.Where("topic_set_level = ?", params.TopicSetLevel)
	}

	if params.TopicType != "" {
		queryFun = queryFun.Where("topic_type = ?", params.TopicType)
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
	Err := queryFun.Model(&topicModel.TopicSet{}).Find(&queryInfo).Count(&res.Data.PagingArgument.Total).Error
	if Err != nil {
		res.Code = e.NO_DATA_EXISTS
		return res
	}

	for _, item := range queryInfo {
		topic := GetTopicData(c, item.TopicIdList)
		helper := TopicInfoReturnData{
			TopicList:          topic,
			ID:                 item.ID,
			Name:               item.Name,
			TopicSetDifficulty: item.TopicSetDifficulty,
			TopicIdList:        item.TopicIdList,
			TopicSetLevel:      item.TopicSetLevel,
			TopicType:          item.TopicType,
			Remark:             item.Remark,
			CreateAt:           item.CreateAt,
		}
		queryInfoReturn = append(queryInfoReturn, helper)
	}

	res.Data.PagingArgument.PageNum = params.PageNum
	res.Data.PagingArgument.PageSize = params.PageSize
	res.Data.Data = queryInfoReturn

	RSetData := res.Data
	Redis.SetValue(queryRedisParams, RSetData, dataRxpirationTime)

	res.Code = e.SUCCESS

	return res
}

func QueryTopicSetSimpleService(c *gin.Context, params QueryTopicSetParams) *queryTopicSetSimpleReturn {
	res := &queryTopicSetSimpleReturn{}
	res.Code = e.SUCCESS
	var queryInfo []topicModel.TopicSet
	var queryInfoReturn []TopicInfoReturnData

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC_SET_SIMPLE)
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &queryInfoReturn)
		if err != nil {
			logging.Debug(err)
		}
		res.Data = queryInfoReturn
		return res
	}

	queryFun := DB.DBLivingExample.Where("is_use = ?", params.IsUse)
	if len(params.Id) > 0 {
		queryFun = queryFun.Where("id IN (?)", params.Id)
	}

	if params.Name != "" {
		queryFun = queryFun.Where("name = ?", params.Name)
	}

	if params.TopicIdList != "" {
		queryFun = queryFun.Where("topic_id_list = ?", params.TopicIdList)
	}

	if params.TopicSetDifficulty != "" {
		queryFun = queryFun.Where("topic_set_difficulty = ?", params.TopicSetDifficulty)
	}

	if params.TopicSetLevel != "" {
		queryFun = queryFun.Where("topic_set_level = ?", params.TopicSetLevel)
	}

	if params.TopicType != "" {
		queryFun = queryFun.Where("topic_type = ?", params.TopicType)
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

	resp := queryFun.Model(&topicModel.TopicSet{}).Find(&queryInfo)
	if resp.Error != nil {
		res.Code = e.NO_DATA_EXISTS
		return res
	}

	for _, item := range queryInfo {
		topic := GetTopicData(c, item.TopicIdList)
		helper := TopicInfoReturnData{
			TopicList:          topic,
			ID:                 item.ID,
			Name:               item.Name,
			TopicSetDifficulty: item.TopicSetDifficulty,
			TopicIdList:        item.TopicIdList,
			TopicSetLevel:      item.TopicSetLevel,
			TopicType:          item.TopicType,
			Remark:             item.Remark,
			CreateAt:           item.CreateAt,
		}
		queryInfoReturn = append(queryInfoReturn, helper)
	}

	res.Data = queryInfoReturn
	Redis.SetValue(queryRedisParams, queryInfo, dataRxpirationTime)

	return res
}

// 获取套题对应的题目数据
func GetTopicData(c *gin.Context, topicIdList string) []topicModel.Topic {

	topId := strings.Split(topicIdList, ",")

	queryParams := topicQuery.QueryTopicSimpleParams{
		Id: topId,
	}
	topicInfo := topicQuery.QueryTopicSimpleService(c, queryParams)

	return topicInfo.Data
}

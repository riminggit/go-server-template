package userAnswerTopicQuery

import (
	"encoding/json"
	"go-server-template/config"
	"go-server-template/model/user"
	"go-server-template/pkg/apiMap"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
	"go-server-template/pkg/utils"
	"go-server-template/src/topicSet/query"
	"strconv"
	"github.com/gin-gonic/gin"
)

func queryUserAnswerTopicService(c *gin.Context, params queryUserAnswerTopicParams) *queryUserAnswerTopicReturn {
	res := &queryUserAnswerTopicReturn{}
	var queryInfo []userModel.UserAnswerTopicRecord
	var RGetData UserAnswerTopicReturnData

	userInfoRes := util.GetUserInfo(c)
	if userInfoRes.Code != e.SUCCESS {
		res.Code = userInfoRes.Code
		return res
	}

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_USER_ANSWER_TOPIC_QUERY)
	userPrefix := userInfoRes.Data.NickName + "/" + strconv.Itoa(userInfoRes.Data.ID) + "/"
	queryRedisParams := userPrefix + interfaceName + string(redisParamsJson)

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

	queryFun := DB.DBLivingExample.Where("is_use = ?", params.IsUse).Where("user_id = ?", userInfoRes.Data.ID)

	if params.TopicSetName != "" {
		queryParams := &topicSetQuery.QueryTopicSetParams{
			Name: params.TopicSetName,
			IsUse: "1",
		}
		resTopicSetName := topicSetQuery.QueryTopicSetSimpleService(c, *queryParams)
		queryFun = queryFun.Where("topic_set_id = ?", resTopicSetName.Data[0].ID)
	}
	
	if params.Id != ""  {
		queryFun = queryFun.Where("id = ?", params.Id)
	}

	if params.TopicIdList != "" {
		queryFun = queryFun.Where("topic_id_list = ?", params.TopicIdList)
	}

	if params.TopicSetId != "" && params.TopicSetName == "" {
		queryFun = queryFun.Where("topic_set_id = ?", params.TopicSetId)
	}

	if params.AnswerNum != "" {
		queryFun = queryFun.Where("answer_num = ?", params.AnswerNum)
	}

	if params.AnswerCorrectNum != "" {
		queryFun = queryFun.Where("answer_correct_num = ?", params.AnswerCorrectNum)
	}

	if params.IsAchieve != "" {
		queryFun = queryFun.Where("is_achieve = ?", params.IsAchieve)
	}

	if params.TopicDifficulty != "" {
		queryFun = queryFun.Where("topic_difficulty = ?", params.TopicDifficulty)
	}

	if params.TopicLevel != "" {
		queryFun = queryFun.Where("topic_level = ?", params.TopicLevel)
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

	if len(params.AnswerStart) > 0 {
		queryFun = queryFun.Where("answer_start between ? and ?", params.AnswerStart[0], params.AnswerStart[1])
	}

	if len(params.AnswerEnd) > 0 {
		queryFun = queryFun.Where("answer_end between ? and ?", params.AnswerEnd[0], params.AnswerEnd[1])
	}

	queryFun = queryFun.Limit(params.PageSize).Offset((params.PageNum - 1) * params.PageSize)

	queryFun.Model(&userModel.UserAnswerTopicRecord{}).Find(&queryInfo).Count(&res.Data.PagingArgument.Total)

	res.Data.PagingArgument.PageNum = params.PageNum
	res.Data.PagingArgument.PageSize = params.PageSize
	res.Data.Data = queryInfo

	RSetData := res.Data
	Redis.SetValue(queryRedisParams, RSetData, dataRxpirationTime)

	res.Code = e.SUCCESS

	return res
}

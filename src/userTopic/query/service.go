package userTopicQuery

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
	"strconv"
	"github.com/gin-gonic/gin"
)

func queryUserTopicService(c *gin.Context, params queryUserTopicParams) *queryUserTopicReturn {
	res := &queryUserTopicReturn{}
	var queryInfo []userModel.UserAddTopic
	var RGetData UserTopicReturnData

	userInfoRes := util.GetUserInfo(c)
	if userInfoRes.Code != e.SUCCESS {
		res.Code = userInfoRes.Code
		return res
	}

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_USER_TOPIC_QUERY)
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
	if params.Id != "" {
		queryFun = queryFun.Where("id = ?", params.Id)
	}

	if params.Title != "" {
		queryFun = queryFun.Where("title = ?", params.Title)
	}

	if params.QuestionType != "" {
		queryFun = queryFun.Where("question_type = ?", params.QuestionType)
	}

	if params.Degree != "" {
		queryFun = queryFun.Where("degree = ?", params.Degree)
	}

	if params.Level != "" {
		queryFun = queryFun.Where("level = ?", params.Level)
	}

	if params.IsBaseTopic != "" {
		queryFun = queryFun.Where("is_base_topic = ?", params.IsBaseTopic)
	}

	if params.IsImportantTopic != "" {
		queryFun = queryFun.Where("is_important_topic = ?", params.IsImportantTopic)
	}

	if params.ComeFrom != "" {
		queryFun = queryFun.Where("come_from = ?", params.ComeFrom)
	}

	if params.ClassifyId != "" {
		queryFun = queryFun.Where("classify_id = ?", params.ClassifyId)
	}

	if params.CompanyId != "" {
		queryFun = queryFun.Where("company_id = ?", params.CompanyId)
	}

	if params.TagId != "" {
		queryFun = queryFun.Where("tag_id = ?", params.TagId)
	}

	if params.TypeId != "" {
		queryFun = queryFun.Where("type_id = ?", params.TypeId)
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

	queryFun.Model(&userModel.UserAddTopic{}).Find(&queryInfo).Count(&res.Data.PagingArgument.Total)

	res.Data.PagingArgument.PageNum = params.PageNum
	res.Data.PagingArgument.PageSize = params.PageSize
	res.Data.Data = queryInfo

	RSetData := res.Data
	Redis.SetValue(queryRedisParams, RSetData, dataRxpirationTime)

	res.Code = e.SUCCESS

	return res
}

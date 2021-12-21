package classifyTypeQuery

import (
	"encoding/json"
	"go-server-template/config"
	"go-server-template/model/type"
	"go-server-template/pkg/e"
	"strconv"

	"github.com/gin-gonic/gin"

	// "go-server-template/pkg/app"
	"go-server-template/model/classify"
	DB "go-server-template/pkg/db"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
)

func QueryTypeService(c *gin.Context, params QueryTypeParams) *QueryReturn {
	res := &QueryReturn{}

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := "query-type:"
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME

	var queryInfo []typeModel.Type

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

	if params.ClassifyName != "" {
		classifyInfo := &classifyModel.Classify{}
		// 查classify中符合的第一条
		DB.DBLivingExample.Where("classify_name = ?", params.ClassifyName).Model(&classifyModel.Classify{}).First(&classifyInfo)
		queryFun = queryFun.Where("classify_id = ?", classifyInfo.ID).Model(&typeModel.Type{}).Find(&queryInfo)
		Redis.SetValue(queryRedisParams, queryInfo, dataRxpirationTime)
		res.Code = e.SUCCESS
		res.Data = queryInfo
	}

	if params.ClassifyId != "" {
		queryFun = queryFun.Where("classify_id = ?", params.ClassifyId)
	}

	if params.Id != "" {
		queryFun = queryFun.Where("id = ?", params.Id)
	}

	if params.IsUse != "" {
		queryFun = queryFun.Where("is_use = ?", params.IsUse)
	}

	if params.TypeName != "" {
		queryFun = queryFun.Where("type_name = ?", params.TypeName)
	}

	queryFun.Model(&typeModel.Type{}).Find(&queryInfo)

	Redis.SetValue(queryRedisParams, queryInfo, dataRxpirationTime)

	res.Code = e.SUCCESS
	res.Data = queryInfo

	return res
}

func QueryTypeMultipleService(c *gin.Context, params QueryTypeMultipleParams) *QueryReturn {
	res := &QueryReturn{}

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := "query-type-multiple:"
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME

	var queryInfo []typeModel.Type

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

	if len(params.ClassifyName) > 0 {
		var classifyInfo []classifyModel.Classify
		// 查classify中符合的第一条
		DB.DBLivingExample.Where("classify_name IN (?)", params.ClassifyName).Model(&classifyModel.Classify{}).Find(&classifyInfo)

		var classifyIdList []string
		for _, item := range classifyInfo {

			// append 往数组追加元素
			classifyIdList = append(classifyIdList, strconv.Itoa(item.ID))
		}

		queryFun = queryFun.Where("classify_id IN (?)", classifyIdList).Model(&typeModel.Type{}).Find(&queryInfo)
		Redis.SetValue(queryRedisParams, queryInfo, dataRxpirationTime)
		res.Code = e.SUCCESS
		res.Data = queryInfo
	}

	if len(params.ClassifyId) > 0 {
		queryFun = queryFun.Where("classify_id IN (?)", params.ClassifyId)
	}

	if len(params.Id) > 0 {
		queryFun = queryFun.Where("id IN (?)", params.Id)
	}

	if params.IsUse != "" {
		queryFun = queryFun.Where("is_use = ?", params.IsUse)
	}

	if len(params.TypeName) > 0 {
		queryFun = queryFun.Where("type_name IN (?)", params.TypeName)
	}

	queryFun.Model(&typeModel.Type{}).Find(&queryInfo)

	Redis.SetValue(queryRedisParams, queryInfo, dataRxpirationTime)

	res.Code = e.SUCCESS
	res.Data = queryInfo

	return res
}

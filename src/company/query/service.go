package companyQuery

import (
	"encoding/json"
	"go-server-template/config"
	"go-server-template/model/company"
	"go-server-template/pkg/apiMap"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
	"github.com/gin-gonic/gin"
)

func QueryCompanyService(c *gin.Context, params QueryCompanyParams) *QueryCompanyReturn {
	res := &QueryCompanyReturn{}

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.GET_QUERY_COMPANY)
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	var queryInfo []companyModel.Company

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
	if params.Id != "" {
		queryFun = queryFun.Where("id = ?", params.Id)
	}

	if params.CompanyName != "" {
		queryFun = queryFun.Where("company_name = ?", params.CompanyName)
	}

	queryFun.Model(&companyModel.Company{}).Find(&queryInfo)

	Redis.SetValue(queryRedisParams, queryInfo, dataRxpirationTime)
	res.Code = e.SUCCESS
	res.Data = queryInfo

	return res
}

func QueryCompanyMultipleService(c *gin.Context, params QueryCompanyMultipleParams) *QueryCompanyReturn {
	res := &QueryCompanyReturn{}

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_COMPANY)
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	var queryInfo []companyModel.Company

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

	if len(params.CompanyName) > 0 {
		queryFun = queryFun.Where("company_name IN (?)", params.CompanyName)
	}

	queryFun.Model(&companyModel.Company{}).Find(&queryInfo)

	Redis.SetValue(queryRedisParams, queryInfo, dataRxpirationTime)
	res.Code = e.SUCCESS
	res.Data = queryInfo

	return res
}

package companyQuery

import (
	"encoding/json"
	"go-server-template/config"
	"go-server-template/model/company"
	"github.com/gin-gonic/gin"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
)

func queryCompanyService(c *gin.Context, params queryCompanyParams) *queryCompanyReturn {
	res := &queryCompanyReturn{}

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := "query-company:"
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
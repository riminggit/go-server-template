package classify

import (
	// "encoding/json"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-server-template/config"
	"go-server-template/model/classify"
	// "go-server-template/pkg/app"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
)

func queryClassifyService(c *gin.Context, params queryClassifyParams) *queryClassifyReturn {
	res := &queryClassifyReturn{}

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := "query-classify:"
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	var classifyInfo []classifyModel.Classify

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &classifyInfo)
		if err != nil {
			logging.Debug(err)
		}
		res.Code = e.SUCCESS
		res.Data = classifyInfo
		return res
	}

	queryFun := DB.DBLivingExample.Where("is_use = ?", params.IsUse)
	if params.Id != "" {
		queryFun = queryFun.Where("id = ?", params.Id)
	}

	if params.ClassifyName != "" {
		queryFun = queryFun.Where("classify_name = ?", params.ClassifyName)
	}

	if params.Rank != "" {
		queryFun = queryFun.Where("rank = ?", params.Rank)
	}

	queryFun.Model(&classifyModel.Classify{}).Find(&classifyInfo)

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	Redis.SetValue(queryRedisParams, classifyInfo, dataRxpirationTime)
	res.Code = e.SUCCESS
	res.Data = classifyInfo

	return res
}

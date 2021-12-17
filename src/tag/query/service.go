package tagQuery

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"go-server-template/config"
	"go-server-template/model/tag"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
)

func queryTagService(c *gin.Context, params queryTagParams) *queryTagReturn {
	res := &queryTagReturn{}

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := "query-tag:"
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	var queryInfo []tagModel.Tag

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

	if params.TagName != "" {
		queryFun = queryFun.Where("tag_name = ?", params.TagName)
	}

	queryFun.Model(&tagModel.Tag{}).Find(&queryInfo)

	Redis.SetValue(queryRedisParams, queryInfo, dataRxpirationTime)
	res.Code = e.SUCCESS
	res.Data = queryInfo

	return res
}

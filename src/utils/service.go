package utils

import (
	"encoding/json"
	projectConfig "go-server-template/config"
	"go-server-template/model/utils"
	"go-server-template/pkg/apiMap"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"

	"github.com/gin-gonic/gin"
)

func AddColorSerices(c *gin.Context, params AddParams) *CreateReturn {
	res := &CreateReturn{}
	res.Code = e.SUCCESS
	res.Msg = "新增成功"

	colorAdd := []utils.Color{}
	for _, item := range params.Color {
		colorAdd = append(colorAdd, utils.Color{Color: item, IsUse: 1})
	}

	CreateFun := DB.DBLivingExample.Model(&utils.Color{}).Create(&colorAdd)
	if CreateFun.Error != nil {
		res.Code = e.CREATE_DATA_FALE
		res.Msg = "新增失败"
	}

	return res
}

func QueryColorService(c *gin.Context) *queryColorReturn {
	res := &queryColorReturn{}
	res.Code = e.SUCCESS
	var RGetData []string
	queryInfo := []utils.Color{}

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_COLOR_QUERY)
	queryRedisParams := interfaceName

	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &RGetData)
		if err != nil {
			logging.Debug(err)
		}
		res.Data = RGetData
		return res
	}

	DB.DBLivingExample.Where("is_use = ?", 1).Model(&utils.Color{}).Find(&queryInfo)

	for _, item := range queryInfo {
		RGetData = append(RGetData, item.Color)
	}

	res.Data = RGetData
	Redis.SetValue(queryRedisParams, RGetData, dataRxpirationTime)

	return res
}

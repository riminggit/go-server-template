package util

import (
	"encoding/json"
	// "fmt"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
	// "reflect"
)

type getRedisDataReturn struct {
	ReturnCode int
	RedisData  interface{}
}

func GetRedisMapData(params interface{}, IName string) *getRedisDataReturn {
	res := &getRedisDataReturn{}

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := IName
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)
	res.RedisData = redisData

	var getData interface{}

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), getData)
		if err != nil {
			logging.Debug(err)
		}
		res.ReturnCode = e.SUCCESS
		res.RedisData = getData
		return res
	}

	return res
}

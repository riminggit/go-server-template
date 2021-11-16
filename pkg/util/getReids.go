package util

import (
	"encoding/json"
	Redis "go-server-template/pkg/redis"
	logging "go-server-template/pkg/log"
	"go-server-template/pkg/e"
)

type GueryParams struct {
	Params        interface{} 
	InterfaceName string
}

type getRedisDataReturn struct {
	ReturnCode int
	RedisData  interface{}
}

func GetRedisMapData(params *GueryParams) *getRedisDataReturn {
	res := &getRedisDataReturn{}

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := "query-type:"
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	res.RedisData = redisData

	var queryInfo interface{}

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &queryInfo)
		if err != nil {
			logging.Debug(err)
		}
		res.ReturnCode = e.SUCCESS
		res.RedisData = queryInfo
		return res
	}

	return res
}

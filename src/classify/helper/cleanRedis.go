package classifyHelper

import (
	Redis "go-server-template/pkg/redis"
	"go-server-template/pkg/apiMap"
)

func CleanRedisQuery()  {

	query1 := apiMap.GetRedisPrefixName(apiMap.GET_QUERY_CLASSIFY)
	query2 := apiMap.GetRedisPrefixName(apiMap.GET_QUERY_CLASSIFY_TYPE)
	query3 := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_CLASSIFY)
	query4 := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_CLASSIFY_TYPE)
	Redis.PrefixDel(query1)
	Redis.PrefixDel(query2)
	Redis.PrefixDel(query3)
	Redis.PrefixDel(query4)
}



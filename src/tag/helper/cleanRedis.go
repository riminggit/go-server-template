package tagHelper

import (
	Redis "go-server-template/pkg/redis"
	"go-server-template/pkg/apiMap"
)

func CleanRedisQuery()  {
	query1 := apiMap.GetRedisPrefixName(apiMap.GET_QUERY_TAG)
	query2 := apiMap.GetRedisPrefixName(apiMap.GET_QUERY_TAG)
	Redis.PrefixDel(query1)
	Redis.PrefixDel(query2)
}



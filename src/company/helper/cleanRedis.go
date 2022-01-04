package companyHelper

import (
	Redis "go-server-template/pkg/redis"
	"go-server-template/pkg/apiMap"
)

func CleanRedisQuery()  {
	query1 := apiMap.GetRedisPrefixName(apiMap.GET_QUERY_COMPANY)
	query2 := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_COMPANY)
	Redis.PrefixDel(query1)
	Redis.PrefixDel(query2)
}


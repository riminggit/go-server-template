package thisHelper

import (
	Redis "go-server-template/pkg/redis"
	"go-server-template/pkg/apiMap"
)

func CleanRedisQuery()  {
	query1 := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC_TAG_MID)
	query2 := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC_TAG_MID_PADING)
	Redis.PrefixDel(query1)
	Redis.PrefixDel(query2)
}



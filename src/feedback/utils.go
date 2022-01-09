package feedback

import (
	"go-server-template/pkg/apiMap"
	Redis "go-server-template/pkg/redis"
	"go-server-template/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CleanRedisQuery(c *gin.Context)  {
	userPrefix := util.GetUserPrefix(c).UserPrefix
	query1 := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_FEEDBACK)
	query2 := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_FEEDBACK_USER) + userPrefix
	Redis.PrefixDel(query1)
	Redis.PrefixDel(query2)
}


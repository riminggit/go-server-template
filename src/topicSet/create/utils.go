package topicSetCreate

import (
	"go-server-template/pkg/e"
	"go-server-template/src/topic/query"
	"math"
	"github.com/gin-gonic/gin"
)

// 查询整体题目信息后计算
func TopSetDifficultyAndLevelCompute(c *gin.Context, topicSetIdList []string) *ComputeReturn {
	res := &ComputeReturn{}
	res.Code = e.SUCCESS

	queryParams := topicQuery.QueryTopicSimpleParams{
		Id: topicSetIdList,
	}
	topicInfo := topicQuery.QueryTopicSimpleService(c, queryParams)

	if topicInfo.Code != e.SUCCESS {
		res.Code = topicInfo.Code
		return res
	}

	diff := 0
	level := 0
	for _, item := range topicInfo.Data {
		diff = diff + (item.Degree + 1) // 简单为0不好计算，往上 +1
		level = level + item.Level
	}

	// + 0.5向下取整，四舍五入
	res.Data.Difficulty = int(math.Floor(float64(diff)/float64(len(topicInfo.Data)) + 0.5)) - 1
	res.Data.Level = int(math.Floor(float64(level)/float64(len(topicInfo.Data)) + 0.5))

	return res
}

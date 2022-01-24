package topicQuery

import (
	topicModel "go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"
	"go-server-template/src/sections"

	"github.com/gin-gonic/gin"
)

func QueryTopicRandomService(c *gin.Context, params QueryTopicRandomParams) *QueryTopicRandomReturn {
	res := &QueryTopicRandomReturn{}
	res.Code = e.SUCCESS
	var queryInfo topicModel.Topic

	queryCount := "1"
	tableName := `topic`

	queryFun := DB.DBLivingExample

	sqlCommon := util.RandSqlCommon(tableName, queryCount)
	if params.IsQueryByUserLevel == "1" {
		sqlHelper := sections.JudgeQueryCondition(c, queryCount, tableName, "level")
		if sqlHelper != "" {
			queryFun.Raw(sqlHelper).Scan(&queryInfo)
		} else {
			queryFun.Raw(sqlCommon).Scan(&queryInfo)
		}
	} else {
		queryFun.Raw(sqlCommon).Scan(&queryInfo)
	}

	if queryFun.Error != nil {
		res.Code = e.NO_DATA_EXISTS
		return res
	}
	if queryInfo.ID == 0 {
		queryFun.Raw(sqlCommon).Scan(&queryInfo)
	}

	res.Data = queryInfo

	return res
}

// func randData(c *gin.Context) topicModel.Topic {
// 	var data topicModel.Topic
// 	count := QueryTopicAllCount(c)
// 	minId := QueryTopicIDMin(c)
// 	getID := util.RandNumC(int(count), minId)
// 	DB.DBLivingExample.Model(&topicModel.Topic{}).Where("is_use = ?", 1).Where("id = ?", getID).Scan(&data)
// 	return data
// }

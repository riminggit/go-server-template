package topicQuery

import (
	"github.com/gin-gonic/gin"
	topicModel "go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"
	"go-server-template/src/sections"
)

func QueryTopicRandomService(c *gin.Context, params QueryTopicRandomParams) *QueryTopicRandomReturn {
	res := &QueryTopicRandomReturn{}
	res.Code = e.SUCCESS
	var queryInfo topicModel.Topic

	queryCount := "1"
	tableName := `topic`

	queryFun := DB.DBLivingExample

	if params.IsQueryByUserLevel == "1" {
		sqlHelper := sections.JudgeQueryCondition(c, queryCount, tableName, "level")
		if sqlHelper != "" {
			queryFun = queryFun.Raw(sqlHelper).Scan(&queryInfo)
		} else {
			queryInfo = randDataFast(c)
		}
	} else {
		queryInfo = randDataFast(c)
	}

	if queryFun.Error != nil {
		res.Code = e.NO_DATA_EXISTS
		return res
	}
	if queryInfo.ID == 0 {
		sqlHelper := util.RandSqlCommon(tableName, queryCount)
		DB.DBLivingExample.Raw(sqlHelper).Scan(&queryInfo)
	}

	res.Data = queryInfo

	return res
}

func randData(c *gin.Context) topicModel.Topic {
	var data topicModel.Topic
	count := QueryTopicAllCount(c)
	minId := QueryTopicIDMin(c)
	getID := util.RandNumC(int(count), minId)
	sqlHelper := "SELECT * FROM topic WHERE is_use = 1 AND id = ?"
	DB.DBLivingExample.Raw(sqlHelper, getID).Scan(&data)
	return data
}

func randDataFast(c *gin.Context) topicModel.Topic {
	var data topicModel.Topic
	sqlHelper := "SELECT*FROM `topic` AS t1"
	sqlMin := "SELECT MIN(id) FROM `topic`"
	sqlMax := "SELECT MAX(id) FROM `topic`"
	sqlHelper2 := "SELECT ROUND(RAND()*((" + sqlMax + ")-(" + sqlMin + "))+(" + sqlMin + ")) AS id)AS t2"
	sqlWhere := "WHERE t1.id>=t2.id"
	DB.DBLivingExample.Raw(sqlHelper).Joins(sqlHelper2).Raw(sqlWhere).Order("t1.id").Limit(1).Scan(&data)
	return data
}

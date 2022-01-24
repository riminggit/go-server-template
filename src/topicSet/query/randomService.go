package topicSetQuery

// 随机查询套题 ,根据参数决定是否根据用户等级来查询套题
import (
	"fmt"
	topicModel "go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"
	"go-server-template/src/sections"

	"github.com/gin-gonic/gin"
)

func QueryTopicSetRandomService(c *gin.Context, params QueryTopicSetRandomParams) *QueryTopicSetRandomReturn {
	res := &QueryTopicSetRandomReturn{}
	queryCount := "1"
	res.Code = e.SUCCESS
	var queryInfo topicModel.TopicSet
	tableName := `topic_set`

	queryFun := DB.DBLivingExample

	sqlCommon := util.RandSqlCommon(tableName, queryCount)
	if params.IsQueryByUserLevel == "1" {
		sqlHelper := sections.JudgeQueryCondition(c, queryCount, tableName, "topic_set_level")
		fmt.Println(sqlHelper, "sqlHelper")
		if sqlHelper != "" {
			queryFun = queryFun.Raw(sqlHelper).Scan(&queryInfo)
		} else {
			queryInfo = randData(c)
		}
	} else {
		queryInfo = randData(c)
	}

	if queryFun.Error != nil {
		res.Code = e.NO_DATA_EXISTS
		return res
	}
	if queryInfo.ID == 0 {
		DB.DBLivingExample.Raw(sqlCommon).Scan(&queryInfo)
	}

	topic := GetTopicData(c, queryInfo.TopicIdList)

	res.Data = TopicInfoReturnData{
		TopicList:          topic,
		ID:                 queryInfo.ID,
		Name:               queryInfo.Name,
		TopicSetDifficulty: queryInfo.TopicSetDifficulty,
		TopicIdList:        queryInfo.TopicIdList,
		TopicSetLevel:      queryInfo.TopicSetLevel,
		TopicType:          queryInfo.TopicType,
		Remark:             queryInfo.Remark,
		CreateAt:           queryInfo.CreateAt,
	}

	return res
}

func randData(c *gin.Context) topicModel.TopicSet {
	var data topicModel.TopicSet
	sqlHelper:= "SELECT*FROM `topic_set` AS t1"
	sqlMin := "SELECT MIN(id) FROM `topic_set`"
	sqlMax := "SELECT MAX(id) FROM `topic_set`"
	sqlHelper2 := "SELECT ROUND(RAND()*((" + sqlMax + ")-(" + sqlMin + "))+(" + sqlMin + ")) AS id)AS t2"
	sqlWhere := "WHERE t1.id>=t2.id"
	DB.DBLivingExample.Raw(sqlHelper).Joins(sqlHelper2).Raw(sqlWhere).Order("t1.id").Limit(1).Scan(&data)
	return data
}

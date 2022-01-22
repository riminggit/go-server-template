package topicSetQuery

// 随机查询套题 ,根据参数决定是否根据用户等级来查询套题
import (
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
	var querySql string
	tableName := `topic_set`

	querySqlHelper := util.RandSql(tableName, queryCount)

	if params.IsQueryByUserLevel == "1" {
		sqlHelper := sections.JudgeQueryCondition(c, queryCount, tableName, "topic_set_level")
		if sqlHelper != "" {
			querySql = sqlHelper
		} else {
			querySql = querySqlHelper
		}
	} else {
		querySql = querySqlHelper
	}

	queryFun := DB.DBLivingExample.Raw(querySql).Scan(&queryInfo)

	if queryFun.Error != nil {
		res.Code = e.NO_DATA_EXISTS
		return res
	}
	if queryInfo.ID == 0 {
		querySql = querySqlHelper
		queryFunAgain := DB.DBLivingExample.Raw(querySql).Scan(&queryInfo)
		if queryFunAgain.Error != nil {
			res.Code = e.NO_DATA_EXISTS
			return res
		}
	}

	topic := GetTopicData(c, queryInfo.TopicSetIdList)

	res.Data = TopicInfoReturnData{
		TopicList:          topic,
		ID:                 queryInfo.ID,
		Name:               queryInfo.Name,
		TopicSetDifficulty: queryInfo.TopicSetDifficulty,
		TopicSetIdList:     queryInfo.TopicSetIdList,
		TopicSetLevel:      queryInfo.TopicSetLevel,
		TopicType:          queryInfo.TopicType,
		Remark:             queryInfo.Remark,
		CreateAt:           queryInfo.CreateAt,
	}

	return res
}

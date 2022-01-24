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
	tableName := `topic_set`

	queryFun := DB.DBLivingExample

	sqlCommon := util.RandSqlCommon(tableName, queryCount)
	if params.IsQueryByUserLevel == "1" {
		sqlHelper := sections.JudgeQueryCondition(c, queryCount, tableName, "topic_set_level")
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
